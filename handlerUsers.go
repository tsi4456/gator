package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("Too many arguments")
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	var currentTag string
	for _, u := range users {
		if u.Name == s.cfg.CurrentUserName {
			currentTag = " (current)"
		} else {
			currentTag = ""
		}
		fmt.Printf("* %s%s\n", u.Name, currentTag)
	}

	return nil
}
