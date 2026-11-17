package main

import (
	"errors"
	"fmt"
	"context"
	"time"
	"github.com/google/uuid"
	"github.com/squidjesus/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("agg command takes no args, 0 expected")
	}

	url := "https://www.wagslane.dev/index.xml"

	feed, error := fetchFeed(context.Background(), url)
	if error != nil {
		return error
	}

	fmt.Printf("%+v\n", feed)

	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) > 2 || len(cmd.args) <= 1 {
		return errors.New("Requires 2 arguments: <name> <url>")
	}

	params := database.AddFeedParams{
		ID: uuid.New(),
		UserID: s.config.Current_user.ID,
		Name: cmd.args[0],
		Url: cmd.args[1],
	}

	feed, err := s.db.AddFeed(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Println("Added feed")
	fmt.Printf("ID: %v\n", feed.ID)
	fmt.Println("Created at: " + feed.CreatedAt.Format(time.DateTime))
	fmt.Println("Updated at: " + feed.UpdatedAt.Format(time.DateTime))
	fmt.Println("Name: " + feed.Name)
	fmt.Println("Url: " + feed.Url)
	fmt.Printf("User ID: %v\n", feed.UserID)
	return nil
}
