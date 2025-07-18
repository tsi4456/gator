package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tsi4456/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Missing name")
	}
	name := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: name})
	if err != nil {
		return err
	}
	s.cfg.SetUser(user.Name)
	fmt.Printf("User %s created!\n", user.Name)
	return nil
}
