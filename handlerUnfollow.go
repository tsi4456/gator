package main

import (
	"context"
	"fmt"

	"github.com/tsi4456/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Incorrect arguments: requires URL only")
	}
	url := cmd.Args[0]

	feed, err := s.db.GetFeedForURL(context.Background(), url)
	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		return err
	}

	return nil
}
