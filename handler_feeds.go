package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/AhGr3y/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, dbUser database.User) error {
	if len(cmd.args) < 2 {
		return errors.New("missing arguments: Usage: gator addFeed <feed name> <feed url>")
	}

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    dbUser.ID,
	}

	dbFeed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	createFeedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    dbUser.ID,
		FeedID:    dbFeed.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), createFeedFollowParams)
	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}

	fmt.Println("Feed created successfully!")
	printFeed(dbFeed, dbUser.Name)

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	dbFeeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching feeds from database: %w", err)
	}

	for _, dbFeed := range dbFeeds {
		dbUser, err := s.db.GetUserByID(context.Background(), dbFeed.UserID)
		if err != nil {
			return fmt.Errorf("error fetching user from database: %w", err)
		}
		printFeed(dbFeed, dbUser.Name)
		println("-----------------------")
	}
	return nil
}

func printFeed(feed database.Feed, userName string) {
	fmt.Printf(" * ID:           %s\n", feed.ID)
	fmt.Printf(" * CreatedAt:    %v\n", feed.CreatedAt)
	fmt.Printf(" * UpdatedAt:    %v\n", feed.UpdatedAt)
	fmt.Printf(" * Name:         %s\n", feed.Name)
	fmt.Printf(" * URL:          %s\n", feed.Url)
	fmt.Printf(" * UserID:       %s\n", feed.ID)
	fmt.Printf(" * UserName:     %s\n", userName)
}
