package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("Too many arguments")
	}

	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return errors.New("failed to delete users")
	}
	fmt.Println("Users deleted!")
	return nil
}
