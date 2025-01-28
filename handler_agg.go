package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/tierant5/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("%v: expects 'time_between_reqs'", cmd.name)
	}
	time_between_reqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("%v: Collecting feeds every %v\n", cmd.name, time_between_reqs)
	ticker := time.NewTicker(time_between_reqs)
	for range ticker.C {
		scrapeFeeds(s, cmd)
	}
	return nil
}

func scrapeFeeds(s *state, cmd command) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID: feed.ID,
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		return err
	}
	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}
	rssFeed.PrintFeed(cmd)
	return nil
}
