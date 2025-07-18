package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	// feedURL := "https://www.wagslane.dev/index.xml"

	// feed, err := fetchFeed(context.Background(), feedURL)
	// if err != nil {
	// 	return err
	// }

	if len(cmd.Args) != 1 {
		return fmt.Errorf("Incorrect arguments: provide only a period for RSS sweeps")
	}

	period, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %s\n", period)
	ticker := time.NewTicker(period)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
