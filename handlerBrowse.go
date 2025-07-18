package main

import (
	"context"
	"fmt"
	"strconv"

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
