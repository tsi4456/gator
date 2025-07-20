package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/tsi4456/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var postLimitString string
	if len(cmd.Args) > 1 {
		return fmt.Errorf("Unwanted arguments provided; command requires only an optional post limit (default: 2)")
	} else if len(cmd.Args) == 1 {
		postLimitString = cmd.Args[0]
	}
	postLimit, err := strconv.Atoi(postLimitString)
	if err != nil {
		postLimit = 2
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{UserID: user.ID, Limit: int32(postLimit)})
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("title: %s\n", post.Title)
		fmt.Printf("published: %s\n", post.PublishedAt)
		fmt.Printf("description: %s\n", post.Description.String)
		fmt.Printf("feed: %s\n\n", post.FeedName)
	}
	return nil
}

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
