package main

import (
	"fmt"
)

type command struct {
	Name      string
	Arguments []string
}

type commands struct {
	Handlers map[string]func(*state, command) error
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
