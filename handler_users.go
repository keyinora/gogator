package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err == nil {
		for i := 0; len(users) > i; i++ {
			fmt.Printf("* %v", users[i].Name)
			if users[i].Name == s.cfg.CurrentUserName {
				fmt.Printf(" (current)")
			}
			fmt.Println()
		}
	}
	return err
}
