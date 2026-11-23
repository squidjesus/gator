package main

import (
	"errors"
	"strconv"
	"fmt"
	"context"
	"time"
	"github.com/google/uuid"
	"github.com/squidjesus/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) < 1 || len(cmd.args) > 1 {
		return errors.New("agg command takes 1 arg, <duration> ex. '1s', '1m', '1h'")
	}

	interval, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %s\n", cmd.args[0])

	ticker := time.NewTicker(interval)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	return nil
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) > 2 || len(cmd.args) <= 1 {
		return errors.New("Requires 2 arguments: <name> <url>")
	}

	params := database.AddFeedParams{
		ID: uuid.New(),
		UserID: user.ID,
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
	fmt.Printf("User ID: %v\n", user.ID)

	// Add follow
	followParams := database.CreateFeedFollowParams{
		ID:			uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		UserID:		user.ID,
		FeedID:		feed.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		return err
	}

	return nil
}


func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("feeds command takes no args, 0 expected")
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println("Name: " + feed.Name)
		fmt.Println("URL: " + feed.Url)
		fmt.Println("Added by: " + feed.Uname)
	}
	return nil
}


func handlerBrowse(s *state, cmd command, user database.User) error {
	postLimit := int32(2)
	if len(cmd.args) > 1 {
		return errors.New("browse command takes 1 optional arg, <number of posts>")
	}
	if len(cmd.args) == 1 {
		num, err := strconv.ParseInt(cmd.args[0], 10, 32)
		if err != nil {
			return err
		}
		postLimit = int32(num)
	}
	param := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: postLimit,
	}
	posts, err := s.db.GetPostsForUser(context.Background(), param)
	if err != nil {
		return err
	}
	for _, post := range posts {
		fmt.Println("Title: " + post.Title)
		fmt.Printf("Published: %v\n", post.PublishedAt)
		fmt.Printf("Description: %s\n", post.Description)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println()
	}

	return nil
}
