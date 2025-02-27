package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	dbUsers, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	if len(dbUsers) == 0 {
		fmt.Println("There are no users in the database.")
		return nil
	}

	currentUser := s.config.CurrentUserName
	for _, dbUser := range dbUsers {
		if dbUser.Name == currentUser {
			fmt.Printf("* %s (current)\n", dbUser.Name)
			continue
		}
		fmt.Printf("* %s\n", dbUser.Name)
	}

	return nil
}
