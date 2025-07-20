package main

import (
	"context"
	"database/sql"
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

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("Incorrect arguments: provide only a period for RSS sweeps")
	}

	period, err := time.ParseDuration("5m")
	if len(cmd.Args) == 1 {
		period, err = time.ParseDuration(cmd.Args[0])
	}
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %s\n", period)
	ticker := time.NewTicker(period)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

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

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return err
	}

	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	for _, feedItem := range feed.Channel.Item {
		pDate, err := time.Parse("Mon, 2 Jan 2006 15:04:05 -0700", feedItem.PubDate)
		if err != nil {
			return err
		}
		nullDesc := sql.NullString{String: feedItem.Description, Valid: true}
		post, err := s.db.CreatePost(context.Background(), database.CreatePostParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Title: feedItem.Title, Url: feedItem.Link, Description: nullDesc, PublishedAt: pDate, FeedID: nextFeed.ID})
		if err != nil {
			return err
		}
		fmt.Println(post.Title)
		fmt.Println(post.PublishedAt)
	}
	return nil
}
