package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/tsi4456/gator/internal/database"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}
	req.Header.Set("User-Agent", "gator")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}

	rawData, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, nil
	}
	defer resp.Body.Close()

	var feed RSSFeed
	if err = xml.Unmarshal(rawData, &feed); err != nil {
		return &RSSFeed{}, nil
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for n, i := range feed.Channel.Item {
		feed.Channel.Item[n].Title = html.UnescapeString(i.Title)
		feed.Channel.Item[n].Description = html.UnescapeString(i.Description)
	}
	return &feed, nil
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
