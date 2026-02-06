package main

import (
	"log"
	"os"

	"github.com/fatbotgw/gator/internal/commands"
	"github.com/fatbotgw/gator/internal/config"
)

func main()  {
	gatorConfig := config.Read()

	progState := &commands.State{
		Cfg: gatorConfig,
	}
	comMap := commands.Commands{
		Handlers: make(map[string]func(*commands.State, commands.Command) error),
	}

	comMap.Register("login", commands.HandlerLogin)
	
	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}

	comName := os.Args[1]
	comArgs := os.Args[2:]

	cmd := commands.Command{
		Name: comName,
		Arguments: comArgs,
	}

	if err := comMap.Run(progState, cmd); err != nil {
		log.Fatal(err)
	}
}
