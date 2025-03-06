package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {

	err := s.db.DeleteUsers(context.Background())
	if err == nil {
		fmt.Println("all users deleted")
	}
	return err
}
