package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
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
	fmt.Printf("Getting feed: %v\n", feedURL)
	res, err := http.Get(feedURL)
	if err != nil {
		return &RSSFeed{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return &RSSFeed{}, err
	}
	fmt.Printf("***BODY:\n%v\n", body)

	var feed RSSFeed
	if err = xml.Unmarshal(body, &feed.Channel); err != nil {
		return &RSSFeed{}, err
	}
	fmt.Printf("***Unmarshaled:\n%v\n", feed)
	for _, feedItem := range feed.Channel.Item {
		fmt.Println(feedItem)
	}
	return &feed, nil
	// return &RSSFeed{}, nil
}

// This will be the long-running aggregator server. Initially, it only fetches
// a single feed to ensure the parsing works.
func handlerFeed(s *state, cmd command) error {
	address := "https://www.wagslane.dev/index.xml"

	fetchFeed(context.Background(), address)

	return nil
}
