package main

import (
	"context"
	"encoding/xml"
	"fmt"
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

func (r *RSSFeed) unescapeHtml() {
	r.Channel.Title = html.UnescapeString(r.Channel.Title)
	r.Channel.Description = html.UnescapeString(r.Channel.Description)
	for _, item := range r.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}
}

func (r *RSSFeed) PrintFeed(cmd command) {
	fmt.Printf("%v: RSSFeed.Channel.Title: %v\n", cmd.name, r.Channel.Title)
	fmt.Printf("%v: RSSFeed.Channel.Link: %v\n", cmd.name, r.Channel.Link)
	fmt.Printf("%v: RSSFeed.Channel.Description: %v\n", cmd.name, r.Channel.Description)
	for i, item := range r.Channel.Item {
		fmt.Printf("%v: RSSFeed.Channel.Item[%v].Title: %v\n", cmd.name, i, item.Title)
		fmt.Printf("%v: RSSFeed.Channel.Item[%v].Link: %v\n", cmd.name, i, item.Link)
		fmt.Printf("%v: RSSFeed.Channel.Item[%v].Description: %v\n", cmd.name, i, item.Description)
		fmt.Printf("%v: RSSFeed.Channel.Item[%v].PubDate: %v\n", cmd.name, i, item.PubDate)
	}
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
	if resp.StatusCode > 299 {
		return &RSSFeed{}, fmt.Errorf("bad http request: %v", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer resp.Body.Close()
	var rssFeed RSSFeed
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return &RSSFeed{}, err
	}
	rssFeed.unescapeHtml()
	return &rssFeed, nil
}
