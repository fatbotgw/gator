package main

import (
	"fmt"

	"github.com/fatbotgw/gator/internal/config"
)

func main()  {
	gatorConfig := config.Read()

	fmt.Printf("Config: %v\n", gatorConfig)

}
