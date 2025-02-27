package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	// if len(cmd.args) == 0 {
	// 	return errors.New("missing argument: Usage: gator agg <feed url>")
	// }

	rssFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	fmt.Println(rssFeed)

	return nil
}
