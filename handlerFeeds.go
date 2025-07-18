package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("Unwanted arguments provided")
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("name: %s\n", feed.Name)
		fmt.Printf("url: %s\n", feed.Url)
		fmt.Printf("created by: %s\n\n", feed.Username)
	}
	return nil
}
