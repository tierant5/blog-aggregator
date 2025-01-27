package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tierant5/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("%v: 'name' and 'url' needed", cmd.name)
	}
	name := cmd.args[0]
	url := cmd.args[1]
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("%v: Feed.ID: %v\n", cmd.name, feed.ID)
	fmt.Printf("%v: Feed.CreatedAt: %v\n", cmd.name, feed.CreatedAt)
	fmt.Printf("%v: Feed.UpdatedAt: %v\n", cmd.name, feed.UpdatedAt)
	fmt.Printf("%v: Feed.Name: %v\n", cmd.name, feed.Name)
	fmt.Printf("%v: Feed.Url: %v\n", cmd.name, feed.Url)
	fmt.Printf("%v: Feed.UserID: %v\n", cmd.name, feed.UserID)
	return nil
}
