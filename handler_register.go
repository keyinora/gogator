package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/keyinora/gogator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]

	// exit code 1 if user already exists
	_, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		fmt.Println("User with that name already exists")
		os.Exit(1)
	}

	currentTime := time.Now()

	parms := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      name,
	}

	_, err = s.db.CreateUser(context.Background(), parms)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("the user '%v' was created\n", name)
	return nil
}
