package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tierant5/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("register: expected a name")
	}
	name := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		return fmt.Errorf("register: user '%v' already exists", name)
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("register: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("register: %w", err)
	}

	fmt.Printf("register: user '%v' created!\n", name)
	fmt.Printf("register: DB obj: %v\n", user)

	return nil
}
