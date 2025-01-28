package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
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
