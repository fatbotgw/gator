package main

import (
	"context"
	"fmt"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("ERROR: Command missing arguments.")
	}
	_, err := s.db.GetUser(context.Background(), cmd.Arguments[0])
	if err != nil {
		log.Fatal("user does not exist")
	}

	username := cmd.Arguments[0]
	err = s.cfg.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Println("Current user has been set.")
	return nil
}

