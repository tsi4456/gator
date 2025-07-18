package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tsi4456/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Incorrect arguments: requires URL only")
	}
	url := cmd.Args[0]

	feed, err := s.db.GetFeedForURL(context.Background(), url)
	if err != nil {
		return err
	}

	feed_follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		return err
	}

	fmt.Printf("feed: %s\n", feed_follow.Feedname)
	fmt.Printf("user: %s\n", feed_follow.Username)
	return nil
}
