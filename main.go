package main

import (
	"log"
	"os"

	"github.com/AhGr3y/gator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	config, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}

	programState := &state{
		config: &config,
	}

	commands := commands{
		registeredCommands: map[string]func(*state, command) error{},
	}
	commands.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("missing argument")
	}

	command := command{
		name: args[1],
		args: args[2:],
	}

	if err := commands.run(programState, command); err != nil {
		log.Fatal(err)
	}

}
