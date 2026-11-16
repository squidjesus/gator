package main

import (
	"errors"
	"fmt"
	"context"
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
