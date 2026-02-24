package main

import (
	"context"
	"fmt"
	"time"

	"github.com/fatbotgw/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Arguments) < 2 {
		return fmt.Errorf("ERROR: Command missing arguments")
	}

	feed_ID := uuid.New()
	created := time.Now()
	updated := time.Now()

	newFeed := database.CreateFeedParams{
		ID: 		feed_ID,
		CreatedAt:	created,
		UpdatedAt:	updated,
		Name: 		cmd.Arguments[0],
		Url: 		cmd.Arguments[1],
		UserID: 	user.ID,
	}
	
	feed, err := s.db.CreateFeed(context.Background(), newFeed)
	if err != nil {
		return err
	}
	followedFeed := database.CreateFeedFollowParams{
		ID:		uuid.New(),
		CreatedAt: 	created,
		UpdatedAt: 	updated,
		UserID: 	user.ID,
		FeedID: 	feed_ID,
	}
	feedRow, err := s.db.CreateFeedFollow(context.Background(), followedFeed)
	if err != nil {
		return err
	}
	printFeed(feed)
	fmt.Println("\nFeed followed successfully:")
	fmt.Printf("%v is now following %v\n", feedRow.UserName, feedRow.FeedName)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Println("Adding Feed:")
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}

