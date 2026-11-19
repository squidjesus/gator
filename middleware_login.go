package main

import (
	"github.com/squidjesus/gator/internal/database"
	"context"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {

	loggedIn := func (s *state, cmd command) error {
		currentUser, err := s.db.GetUser(context.Background(), s.config.Current_user_name)
		if err != nil {
			return err
		}
		return handler(s, cmd, currentUser)
	}	
	return loggedIn
}
