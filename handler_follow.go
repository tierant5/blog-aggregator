package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tierant5/gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("%v: expects 'url'", cmd.name)
	}
	url := cmd.args[0]
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("%v: Feed: %v, User: %v\n", cmd.name, feed.Name, user.Name)
	return nil
}
