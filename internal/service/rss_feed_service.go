package service

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/model"
)

func FetchFeed(url string) (model.RSSFeed, error) {
	res, err := http.Get(url)
	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	rssFeed := model.RSSFeed{}
	if err := xml.Unmarshal(body, &rssFeed); err != nil {
		return rssFeed, fmt.Errorf("decode json: %w", err)
	}
	return rssFeed, nil
}

func scrapFeed(feedService *FeedService, postService *PostService, feed model.Feed, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("Fetching from %s...\n", feed.Url)

	_, err := feedService.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error marking feed as fetched. Feed URL: %s, Error: %v", feed.Url, err)
		return
	}

	rssFeed, err := FetchFeed(feed.Url)
	if err != nil {
		log.Printf("Fetching from %s error: %v", feed.Url, err)
		return
	}

	for _, post := range rssFeed.Channel.Item {
		_, err := postService.CreatePost(context.Background(), post.Title, post.Link, post.Description, post.PubDate, feed.ID)
		if err != nil {
			if strings.Contains(err.Error(), "violates unique constraint") {
				continue
			}
			log.Printf("Creating post error: %v", err)
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))

}

func StartScraping(feedService *FeedService, postService *PostService, interval time.Duration) {
	ticker := time.NewTicker(interval)
	wg := &sync.WaitGroup{}
	for ; ; <-ticker.C {
		feedsToFetch, err := feedService.GetNextFeedsToFetch(context.Background(), 5)
		if err != nil {
			log.Println("Cannot get feeds to fetch")
			continue
		}
		wg.Add(len(feedsToFetch))
		for _, feed := range feedsToFetch {
			go scrapFeed(feedService, postService, feed, wg)
		}
		wg.Wait()
	}
}
