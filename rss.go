package main

import (
	"net/http"
	"context"
	"time"
	"io"
	"encoding/xml"
	"html"
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

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var feed RSSFeed
	if err := xml.Unmarshal(body, &feed); err != nil {
		return nil, err
	}

	// Clean up HTML entities
	feed.Channel.Title = (html.UnescapeString(feed.Channel.Title))
	feed.Channel.Description = (html.UnescapeString(feed.Channel.Description))
	for i, _ := range feed.Channel.Item {
		t := feed.Channel.Item[i].Title
		d := feed.Channel.Item[i].Description
		feed.Channel.Item[i].Title = (html.UnescapeString(t))
		feed.Channel.Item[i].Description = (html.UnescapeString(d))
	}

	return &feed, nil
}
