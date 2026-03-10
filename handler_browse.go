package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/fatbotgw/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	postLimit := 2
	var err error
	if len(cmd.Arguments) >= 1 {
		postLimit, err = strconv.Atoi(cmd.Arguments[0])
		if err != nil {
			return err
		}
	}
	userParams := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: int32(postLimit),
	}
	feedPosts, err := s.db.GetPostsForUser(context.Background(), userParams)
	if err != nil {
		return err
	}
	for _, post := range feedPosts {
		fmt.Println(post.Title)
		fmt.Println(post.PublishedAt)
		fmt.Println(post.Description)
	}
	return nil
}

