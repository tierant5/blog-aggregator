package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/tierant5/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) > 0 {
		limit, _ = strconv.Atoi(cmd.args[0])
	}
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}
	for _, post := range posts {
		fmt.Printf("%v: Post.Title: %v\n", cmd.name, post.Title)
	}
	return nil
}
