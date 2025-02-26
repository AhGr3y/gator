package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AhGr3y/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("missing argument: Usage: gator %s <user name>", cmd.name)
	}

	userName := cmd.args[0]
	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      userName,
	}

	dbUser, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.config.SetUser(dbUser.Name); err != nil {
		return fmt.Errorf("error updating config: %w", err)
	}

	fmt.Printf("%s was created successfully!\n", dbUser.Name)
	log.Printf("User created successfully: %v\n", dbUser)

	return nil
}
