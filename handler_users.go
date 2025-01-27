package main

import (
	"context"
	"fmt"
)

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
