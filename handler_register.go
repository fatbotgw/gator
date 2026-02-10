package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fatbotgw/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("ERROR: Command missing arguments.")
	}
	_, err := s.db.GetUser(context.Background(), cmd.Arguments[0])
	if err == nil {
		log.Fatal("user already exists")
	}
	newUser := database.CreateUserParams{
		ID:	uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		Name:		cmd.Arguments[0],
	}
	s.db.CreateUser(context.Background(), newUser)


	err = s.cfg.SetUser(cmd.Arguments[0])
	if err != nil {
		return err
	 }
	fmt.Printf("Current user (%s) has been set.", cmd.Arguments[0])
	return nil
}
