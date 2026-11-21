package main

import (
	"net/http"
	"context"
	"fmt"
	"time"
	"io"
	"encoding/xml"
	"html"
	"github.com/squidjesus/gator/internal/database"
	"database/sql"
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


func scrapeFeeds(s *state) error {
	ctx := context.Background()
	// get next feed to fetch from the DB
	nextFeed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}
	// mark it as fetched
	nowTime := sql.NullTime{
		Time: time.Now(),
		Valid: true,
	}
	param := database.MarkFeedFetchedParams{
		LastFetchedAt: nowTime,
		UpdatedAt:	   time.Now(),
		ID:			   nextFeed.ID,
	}
	if err := s.db.MarkFeedFetched(ctx, param); err != nil {
		return err
	}
	// fetch the feed by URL
	fetchedFeed, err := fetchFeed(ctx, nextFeed.Url)
	if err != nil {
		return err
	}
	fmt.Println("Feed: " + nextFeed.Name + " " + nextFeed.Url)
	// iterate over items and print 'title'
	for _, item := range fetchedFeed.Channel.Item {
		fmt.Println(" * " + item.Title)
	}
	return nil
}

