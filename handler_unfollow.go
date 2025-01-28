package main

import (
	"context"
	"fmt"

	"github.com/tierant5/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("%v: expects 'url'", cmd.name)
	}
	url := cmd.args[0]
	err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url:    url,
	})
	if err != nil {
		return err
	}
	return nil
}
