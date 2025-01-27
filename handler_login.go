package main

import (
	"context"
	"fmt"
)

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
