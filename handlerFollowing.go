package main

import (
	"context"
	"fmt"

	"github.com/tsi4456/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("No arguments expected")
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("feed: %s\n", feed.Feedname)
	}
	return nil
}
