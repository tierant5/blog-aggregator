package main

import (
	"context"
	"fmt"

	"github.com/tierant5/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	fmt.Printf("%v: %v is following:\n", cmd.name, s.cfg.CurrentUserName)
	for _, feed := range feeds {
		fmt.Printf("  - %v\n", feed.FeedName)
	}
	return nil
}
