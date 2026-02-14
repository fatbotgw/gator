package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"html"
	"time"
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

// This will be the long-running aggregator server. Initially, it only fetches
// a single feed to ensure the parsing works.
func handlerFeed(s *state, cmd command) error {
	address := "https://www.wagslane.dev/index.xml"

	res, err := fetchFeed(context.Background(), address)
	if err != nil {
		return err
	}

	fmt.Println(res)
	return nil
}
