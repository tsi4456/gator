package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tsi4456/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Incorrect arguments: requires feed name and URL")
	}
	url := cmd.Args[0]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: user.Name, Url: url, UserID: user.ID})
	if err != nil {
		return err
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		return err
	}

	fmt.Printf("id: %d\n", feed.ID)
	fmt.Printf("created_at: %s\n", feed.CreatedAt)
	fmt.Printf("updated_at: %s\n", feed.UpdatedAt)
	fmt.Printf("name: %s\n", feed.Name)
	fmt.Printf("url: %s\n", feed.Url)
	fmt.Printf("user_id: %s\n", feed.UserID)
	return nil
}
