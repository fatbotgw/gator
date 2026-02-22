package main

import (
	"context"
	"fmt"
	"time"

	"github.com/fatbotgw/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFeed(s *state, cmd command) error {
	if len(cmd.Arguments) < 2 {
		return fmt.Errorf("ERROR: Command missing arguments")
	}
	curUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUser)

	feed_ID := uuid.New()
	created := time.Now()
	updated := time.Now()

	newFeed := database.CreateFeedParams{
		ID: 		feed_ID,
		CreatedAt:	created,
		UpdatedAt:	updated,
		Name: 		cmd.Arguments[0],
		Url: 		cmd.Arguments[1],
		UserID: 	curUser.ID,
	}
	
	feed, err := s.db.CreateFeed(context.Background(), newFeed)
	if err != nil {
		return err
	}
	followedFeed := database.CreateFeedFollowParams{
		ID:		uuid.New(),
		CreatedAt: 	created,
		UpdatedAt: 	updated,
		UserID: 	curUser.ID,
		FeedID: 	feed_ID,
	}
	feedRow, err := s.db.CreateFeedFollow(context.Background(), followedFeed)
	if err != nil {
		return err
	}
	fmt.Printf("%v", feedRow)
	printFeed(feed)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}

