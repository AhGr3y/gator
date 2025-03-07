package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/AhGr3y/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		fmt.Println("invalid argument: Usage: gator agg <refresh interval>")
		return nil
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		fmt.Printf("failed to parse duration: %v", err)
		return nil
	}
	ticker := time.NewTicker(timeBetweenReqs)

	fmt.Printf("Collecting feeds every %v...\n", cmd.args[0])

	for range ticker.C {
		scrapeFeeds(s)
	}

	return nil
}

func scrapeFeeds(s *state) {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Printf("failed to fetch feed: %v", err)
		return
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
		fmt.Printf("failed to mark feed as fetched: %v", err)
		return
	}

	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		fmt.Printf("failed to fetch feed: %v", err)
		return
	}

	fmt.Printf("Found %v post(s) at %v.\n", len(feed.Channel.Item), feed.Channel.Title)
	fmt.Println("Saving posts...")
	for _, item := range feed.Channel.Item {
		// Ignore item if post already saved in database
		_, err := s.db.GetPostByURL(context.Background(), item.Link)
		if err == nil {
			continue
		} else {
			if errors.Is(err, sql.ErrNoRows) {
				// Do nothing
			} else {
				log.Printf("failed to get post by url: %v\n", err)
				continue
			}

		}

		pubDate, err := parseDate(item.PubDate)
		if err != nil {
			log.Printf("failed to parse date: %v\n", err)
			continue
		}

		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: pubDate,
			FeedID:      nextFeed.ID,
		}

		post, err := s.db.CreatePost(context.Background(), params)
		if err != nil {
			log.Printf("failed to create post: %v\n", err)
			continue
		}

		fmt.Printf("%v saved successfully.\n", post.Title)
	}

	fmt.Println("Posts saved successfully.")

}

// parseDate - Parses date string with multiple time layouts until
// the correct layout is found.
func parseDate(date string) (time.Time, error) {
	layouts := []string{
		time.RFC1123Z,
		time.RFC1123,
	}

	for _, layout := range layouts {
		pubDate, err := time.Parse(layout, date)
		if err == nil {
			return pubDate, nil
		}
	}

	return time.Time{}, errors.New("unknown date layout")
}
