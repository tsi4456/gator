package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tsi4456/gator/internal/database"
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
