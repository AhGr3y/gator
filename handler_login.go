package main

import (
	"context"
	"fmt"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("missing argument: Usage: gator %s <user name>", cmd.name)
	}

	userName := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.config.SetUser(userName); err != nil {
		return err
	}

	fmt.Printf("%s has successfully logged in!\n", userName)

	return nil
}
