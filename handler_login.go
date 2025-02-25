package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("missing argument for the command: %s", cmd.name)
	}

	username := cmd.args[0]
	if err := s.config.SetUser(username); err != nil {
		return err
	}

	fmt.Printf("%s has successfully logged in!\n", username)

	return nil
}
