package main

import (
	"fmt"

	"github.com/fatbotgw/gator/internal/config"
)

type state struct {
	Cfg config.Config
}

type command struct {
	Name      string
	Arguments []string
}

type commands struct {
	Handlers map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("ERROR: Command missing arguments.")
	}
	username := cmd.Arguments[0]
	err := s.Cfg.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Println("Current user has been set.")
	return nil
}

// This method registers a new handler function for a command name.
func (c *commands) Register(name string, f func(*state, command) error) {
	c.Handlers[name] = f
}

// This method runs a given command with the provided state if it exists.
func (c *commands) Run(s *state, cmd command) error {
	handler, ok := c.Handlers[cmd.Name]
	if !ok {
		return fmt.Errorf("command not found: %s", cmd.Name)
	}
	return handler(s, cmd)
}
