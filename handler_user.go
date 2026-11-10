package main

import (
	"fmt"
	"errors"
)

func handlerLogin(s *state, cmd command) error {

	if len(cmd.args) > 1 || len(cmd.args) == 0 {
		return errors.New("Requires 1 argument: <username>")
	}

	err := s.config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("%s logged in\n", cmd.args[0])
	return nil
}
