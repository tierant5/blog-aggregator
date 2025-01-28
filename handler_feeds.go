package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("%v: error getting all feed info from the database: %v", cmd.name, err)
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("%v: Feed.ID: %v\n", cmd.name, feed.ID)
		fmt.Printf("%v: Feed.CreatedAt: %v\n", cmd.name, feed.CreatedAt)
		fmt.Printf("%v: Feed.UpdatedAt: %v\n", cmd.name, feed.UpdatedAt)
		fmt.Printf("%v: Feed.Name: %v\n", cmd.name, feed.Name)
		fmt.Printf("%v: Feed.Url: %v\n", cmd.name, feed.Url)
		fmt.Printf("%v: Feed.UserID: %v\n", cmd.name, feed.UserID)
		fmt.Printf("%v: Created By: %v\n", cmd.name, user.Name)
	}

	return nil
}
