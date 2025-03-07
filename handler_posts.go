package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/AhGr3y/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.args) > 1 {
		fmt.Printf("Usage: gator %v [<post limit>]\n", cmd.name)
		return nil
	}

	limit := 2
	if len(cmd.args) == 1 {
		converted, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			fmt.Printf("string conversion failed: %v\n", err)
			return nil
		}
		limit = converted
	}

	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := s.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		fmt.Printf("failed to get posts for user: %v\n", err)
		return nil
	}

	fmt.Printf("Retrieved %v post(s):\n", len(posts))
	for _, post := range posts {
		fmt.Println("")
		fmt.Printf(" * Title: %v\n", post.Title)
		fmt.Printf(" * Description: %v\n", post.Description)
		fmt.Printf(" * Published At: %v\n", post.PublishedAt)
	}

	return nil
}
