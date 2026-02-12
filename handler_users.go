package main

import (
	"context"
	"fmt"
	"log"
)

func handlerUsers(s * state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		log.Fatal("no users in database")
	}
	for i := 0; i < len(users); i++ {
		if s.cfg.CurrentUser == users[i] {
			fmt.Printf("* %s (current)\n", users[i])
			continue
		}
		fmt.Printf("* %s\n", users[i])
	}

	return  nil
}
