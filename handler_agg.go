package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"html"
	"time"

	"github.com/fatbotgw/gator/internal/database"
	"github.com/google/uuid"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return &RSSFeed{}, err
	}

	var feed RSSFeed
	if err = xml.Unmarshal(body, &feed); err != nil {
		return &RSSFeed{}, err
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i := range feed.Channel.Item {
	    feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
	    feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}

	return &feed, nil
}

// This is the long-running aggregator server.
func handlerAgg(s *state, cmd command) error {
	// return scrapeFeeds(s)
	if len(cmd.Arguments) < 1 {
		return fmt.Errorf("ERROR: Command missing arguments")
	}

	time_between_reqs := cmd.Arguments[0]

	timeBetweenRequests, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		fmt.Println("\n*** SCRAPING FEED ***\n")
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	// Get the next feed to fetch from the DB.
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	// Mark it as fetched.
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}
	// Fetch the feed using the URL (we already wrote this function)
	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}
	// Iterate over the items in the feed and print their titles to the console.
	for _, item := range feedData.Channel.Item {
		// Examples of publish dates from RSS scrapes:
		// Mon, 01 Jan 0001 00:00:00 +0000
		// Wed, 10 Mar 2021 00:00:00 +0000
		layout := "Mon, 01 Jan 0001 00:00:00 +0000"

		date, _ := time.Parse(layout, item.PubDate)

		// create a post for each item using the CreatePostParams struct
		// and CreatePost() which inserts into db
		post := database.CreatePostParams{
			ID:			uuid.New(),
			CreatedAt:	time.Now(),
			UpdatedAt:	time.Now(),
			Title:		item.Title,
			Url:		item.Link,
			Description:item.Description,
			PublishedAt:date,
			FeedID:		feed.ID,
		}
		_, err := s.db.CreatePost(context.Background(), post)
		if err != nil {
			return err
		}
	}	

	return nil
}
