package main

import (
	"context"
	"fmt"
	
)

func handlerFeeds(s *state, cmd command) error {
	feedList, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return  err
	}

	for _, feed := range feedList {
		userName, err := s.db.GetUserName(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println(userName)
	}

	return nil
}
