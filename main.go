package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/AhGr3y/gator/internal/config"
	"github.com/AhGr3y/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

func main() {
	config, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}

	db, err := sql.Open("postgres", config.DbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	programState := &state{
		db:     dbQueries,
		config: &config,
	}

	commands := commands{
		registeredCommands: map[string]func(*state, command) error{},
	}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", handlerAddFeed)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("missing argument")
	}

	cmdName := args[1]
	cmdArgs := args[2:]
	command := command{
		name: cmdName,
		args: cmdArgs,
	}

	if err := commands.run(programState, command); err != nil {
		log.Fatal(err)
	}

}
