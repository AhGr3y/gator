package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/AhGr3y/gator/internal/database"
	"github.com/google/uuid"
)

// handlerFollow - takes a single url argument and creates a new feed follow record
// for the current user. It prints the name of the feed and the current user.
func handlerFollow(s *state, cmd command, dbUser database.User) error {
	if len(cmd.args) == 0 {
		return errors.New("missing argument: Usage: gator follow <feed url>")
	}

	dbFeed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("error fetching feed by url: %w", err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    dbUser.ID,
		FeedID:    dbFeed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}

	fmt.Println("Successfully created feed follow!")
	fmt.Printf("Feed Name: %v\n", feedFollow.FeedName)
	fmt.Printf("User Name: %v\n", feedFollow.UserName)

	return nil
}

func handlerFollowing(s *state, cmd command, dbUser database.User) error {
	if len(cmd.args) > 0 {
		return errors.New("command 'following' does not take any arguments")
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), dbUser.ID)
	if err != nil {
		return fmt.Errorf("error fetching feed follows for user: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No record found.")
		return nil
	}

	fmt.Printf("Found %v feed(s):\n", len(feedFollows))
	for _, feedFollow := range feedFollows {
		dbFeed, err := s.db.GetFeedByID(context.Background(), feedFollow.FeedID)
		if err != nil {
			return fmt.Errorf("error fetching feed by id: %w", err)
		}
		fmt.Printf(" * %v\n", dbFeed.Name)
	}

	return nil
}

func handlerUnfollow(s *state, cmd command, dbUser database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("invalid argument. Usage: gator %s <feed url>")
	}

	dbFeed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}

	params := database.DeleteFeedFollowForUserParams{
		UserID: dbUser.ID,
		FeedID: dbFeed.ID,
	}

	if err := s.db.DeleteFeedFollowForUser(context.Background(), params); err != nil {
		return fmt.Errorf("failed to unfollow: %w", err)
	}

	fmt.Printf("Unfollowed '%s' successfully!\n", dbFeed.Name)

	return nil
}
