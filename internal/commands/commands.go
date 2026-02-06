package commands

import (
	"fmt"

	"github.com/fatbotgw/gator/internal/config"
)

type State struct {
	Cfg config.Config
}

type Command struct {
	Name      string
	Arguments []string
}

type Commands struct {
	Handlers map[string]func(*State, Command) error
}

func HandlerLogin(s *State, cmd Command) error {
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
func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Handlers[name] = f
}

// This method runs a given command with the provided state if it exists.
func (c *Commands) Run(s *State, cmd Command) error {
	handler, ok := c.Handlers[cmd.Name]
	if !ok {
		return fmt.Errorf("command not found: %s", cmd.Name)
	}
	return handler(s, cmd)
}
