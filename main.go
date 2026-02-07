package main

import (
	"log"
	"os"

	"github.com/fatbotgw/gator/internal/config"
)

func main()  {
	gatorConfig := config.Read()

	progState := &state{
		Cfg: gatorConfig,
	}
	comMap := commands{
		Handlers: make(map[string]func(*state, command) error),
	}

	comMap.Register("login", handlerLogin)
	
	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}

	comName := os.Args[1]
	comArgs := os.Args[2:]

	cmd := command{
		Name: comName,
		Arguments: comArgs,
	}

	if err := comMap.Run(progState, cmd); err != nil {
		log.Fatal(err)
	}
}
