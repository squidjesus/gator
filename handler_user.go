package main
import (
	"fmt"
	"errors"
	"github.com/squidjesus/gator/internal/database"
	"time"
	"context"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {

	if len(cmd.args) > 1 || len(cmd.args) == 0 {
		return errors.New("Requires 1 argument: <username>")
	}

	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("%s logged in\n", user.Name)
	return nil
}


func handlerRegister( s *state, cmd command) error {

	if len(cmd.args) > 1 || len(cmd.args) == 0 {
		return errors.New("Requires 1 argument: <username>")
	}

	userParams := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.args[0],
	}

	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return err
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("Registered new user")
	fmt.Println("Name: " + user.Name)
	fmt.Printf("ID: %v\n", user.ID)
	fmt.Println("Created at: " + user.CreatedAt.Format(time.DateTime))
	fmt.Println("Updated at: " + user.UpdatedAt.Format(time.DateTime))
	return nil
}


func handlerReset(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("reset command takes no args, 0 expected")
	}
	deleted, err := s.db.ResetUsers(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Users table reset:")
	fmt.Printf("Removed %d entries\n", deleted)
	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("reset command takes no args, 0 expected")
	}
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		if user == s.config.Current_user_name{
			fmt.Printf("* %s (current)\n", user)
			continue
		}
		fmt.Printf("* %s\n", user)
	}
	return nil
}
