package main

import (
	"fmt"

	"github.com/fatbotgw/gator/internal/config"
	"github.com/fatbotgw/gator/internal/commands"
)

func main()  {
	gatorConfig := config.Read()
	// fmt.Printf("Initial Config: %v\n", gatorConfig)

	// write username to config file
	config.SetUser(gatorConfig)

	// read config file again and print to console
	gatorConfig = config.Read()
	fmt.Printf("Config: %v\n", gatorConfig)
}
