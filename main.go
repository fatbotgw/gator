package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/fatbotgw/gator/internal/config"
	"github.com/fatbotgw/gator/internal/database"

	_ "github.com/lib/pq"
)

// type state struct {
// 	cfg config.Config
// }

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main()  {
	gatorConfig := config.Read()

	dbURL := gatorConfig.Database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)


	progState := &state{
		db: dbQueries,
		cfg: &gatorConfig,
	}
	comMap := commands{
		Handlers: make(map[string]func(*state, command) error),
	}

	comMap.Register("login", handlerLogin)
	comMap.Register("register", handlerRegister)
	comMap.Register("reset", handlerReset)
	comMap.Register("users", handlerUsers)
	comMap.Register("agg", handlerAgg)
	comMap.Register("addfeed", handlerFeed)
	comMap.Register("feeds", handlerFeeds)
	
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
