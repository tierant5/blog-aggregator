package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"
	rssFeed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return err
	}
	fmt.Printf("%v: RSSFeed.Channel.Title: %v\n", cmd.name, rssFeed.Channel.Title)
	fmt.Printf("%v: RSSFeed.Channel.Link: %v\n", cmd.name, rssFeed.Channel.Link)
	fmt.Printf("%v: RSSFeed.Channel.Description: %v\n", cmd.name, rssFeed.Channel.Description)
	for i, item := range rssFeed.Channel.Item {
		fmt.Printf("%v: RSSFeed.Channel.Item[%v].Title: %v\n", cmd.name, i, item.Title)
		fmt.Printf("%v: RSSFeed.Channel.Item[%v].Link: %v\n", cmd.name, i, item.Link)
		fmt.Printf("%v: RSSFeed.Channel.Item[%v].Description: %v\n", cmd.name, i, item.Description)
		fmt.Printf("%v: RSSFeed.Channel.Item[%v].PubDate: %v\n", cmd.name, i, item.PubDate)
	}
	return nil
}
