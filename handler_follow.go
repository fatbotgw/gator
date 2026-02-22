package main

import (
	"context"
	"fmt"
	"time"

	"github.com/fatbotgw/gator/internal/database"
	"github.com/google/uuid"
)

func follow(s *state, cmd command) error {
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUser)
	if err != nil {
		return err
	}
	feedName, err := s.db.GetFeedNameByURL(context.Background(), cmd.Arguments[0])
	if err != nil {
		return err
	}
	followedFeed := database.CreateFeedFollowParams{
		ID: 	   uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    feedName.ID,
	}
	followRow, err := s.db.CreateFeedFollow(context.Background(), followedFeed)
	if err != nil {
		return err
	}
	printFollowedFeed(followRow)
	return nil
}

func printFollowedFeed(feed database.CreateFeedFollowRow) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.FeedName)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}

func following(s *state, cmd command) error {
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUser)
	if err != nil {
		return err
	}
	feedList, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return err
	}
	for _, feed := range feedList {
		fmt.Println(feed.FeedName)
	}

	return nil
}
