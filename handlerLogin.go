package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Missing username")
	}
	user, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("no such user")
	}
	if err = s.cfg.SetUser(user.Name); err != nil {
		return fmt.Errorf("unable to set username")
	}
	fmt.Printf("Username set to '%s'\n", user.Name)
	return nil
}
