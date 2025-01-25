package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tierant5/gator/internal/config"
	"github.com/tierant5/gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.cmds[cmd.name]
	if !ok {
		return fmt.Errorf("command: %v not registered", cmd.name)
	}
	err := f(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func handerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("login: expected a username")
	}
	name := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("login: user '%v' doesn't exist", name)
	}
	err = s.cfg.SetUser(name)
	if err != nil {
		return err
	}
	fmt.Printf("login: user '%v' was set\n", s.cfg.CurrentUserName)
	return nil
}

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

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("%v: error reseting the database: %v", cmd.name, err)
	}
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("%v: error getting all user info from the database: %v", cmd.name, err)
	}

	for _, user := range users {
		userStr := fmt.Sprintf("* %v", user.Name)
		if user.Name == s.cfg.CurrentUserName {
			userStr += " (current)"
		}
		fmt.Println(userStr)
	}

	return nil
}
