package main

import (
	"context"
	"fmt"
	"time"

	"github.com/keyinora/gogator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	/*

				name: The name of the feed
		url: The URL of the feed
		At the top of the handler, get the current user from the database and connect the feed to that user.

		If everything goes well, print out the fields of the new feed record.
	*/

	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	// Store arguments
	name := cmd.Args[0]
	url := cmd.Args[1]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("user: %s doesn't exist", s.cfg.CurrentUserName)
	}

	currentTime := time.Now()
	//create db parms struct
	parms := database.CreateFeedParams{
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}
	// Add feed
	feed, err := s.db.CreateFeed(context.Background(), parms)
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	fmt.Printf("Feed created: %+v\n", feed)
	return nil
}
