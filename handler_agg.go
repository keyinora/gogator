package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	// Fetch the feed and print it
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}
	fmt.Printf("%+v\n", feed)
	return nil
}
