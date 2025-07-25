package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
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
