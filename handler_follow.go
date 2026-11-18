package main

import (
	"fmt"
	"errors"
	"github.com/squidjesus/gator/internal/database"
	"time"
	"context"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) > 1 || len(cmd.args) == 0 {
		return errors.New("Requires 1 argument: <url>")
	}
	
	url := cmd.args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	params := database.CreateFeedFollowParams{
		ID:			uuid.New(),
		CreatedAt:  time.Now(),
		UpdatedAt:	time.Now(),
		UserID:		s.config.Current_user.ID,
		FeedID:		feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Println("Feed Follow Added:")
	fmt.Println("Name: " + feedFollow.FeedName)
	fmt.Println("Current User: " + s.config.Current_user.Name)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("following command takes no args, 0 expected")
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), s.config.Current_user.ID)
	if err != nil {
		return err
	}

	fmt.Println("User: " + s.config.Current_user.Name)
	for _, follow := range follows {
		fmt.Printf(" * %s\n", follow.FeedName)
	}

	return nil
}

