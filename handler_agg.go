package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/AhGr3y/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("invalid argument: Usage: gator agg <refresh interval>")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}
	ticker := time.NewTicker(timeBetweenReqs)

	fmt.Printf("Collecting feeds every %v\n", cmd.args[0])

	for range ticker.C {
		if err := scrapeFeeds(s); err != nil {
			return fmt.Errorf("error scraping feeds: %w", err)
		}
	}

	return nil
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}

	params := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
		UpdatedAt: time.Now().UTC(),
		ID:        nextFeed.ID,
	}

	if err := s.db.MarkFeedFetched(context.Background(), params); err != nil {
		return fmt.Errorf("failed to mark feed as fetched: %w", err)
	}

	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}

	fmt.Printf("Found %v post(s) at %v:\n", len(feed.Channel.Item), feed.Channel.Title)
	for _, post := range feed.Channel.Item {
		fmt.Printf(" * %v\n", post.Title)
	}

	return nil
}
