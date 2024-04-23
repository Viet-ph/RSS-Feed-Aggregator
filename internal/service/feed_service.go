package service

import (
	"context"
	"time"

	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/database"
	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/model"
	"github.com/google/uuid"
)

type FeedService struct {
	Queries *database.Queries
}

func NewFeedService(q *database.Queries) *FeedService {
	return &FeedService{
		Queries: q,
	}
}

func (feedService *FeedService) CreateFeed(ctx context.Context, name, url string, userId uuid.UUID) (model.Feed, error) {
	feed, err := feedService.Queries.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    userId,
	})

	if err != nil {
		return model.Feed{}, err
	}

	return model.DbFeedToFeed(&feed), nil
}

func (feedService *FeedService) GetFeeds(ctx context.Context) ([]model.Feed, error) {
	dbFeeds, err := feedService.Queries.GetFeeds(ctx)
	if err != nil {
		return nil, err
	}
	feeds := make([]model.Feed, 0, len(dbFeeds))
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, model.DbFeedToFeed(&dbFeed))
	}

	return feeds, nil
}

func (feedService *FeedService) GetNextFeedsToFetch(ctx context.Context, limit int32) ([]model.Feed, error) {
	dbFeeds, err := feedService.Queries.GetNextFeedsToFetch(ctx, limit)
	if err != nil {
		return nil, err
	}
	feeds := make([]model.Feed, 0, len(dbFeeds))
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, model.DbFeedToFeed(&dbFeed))
	}

	return feeds, nil
}

func (feedService *FeedService) MarkFeedFetched(ctx context.Context, feedId uuid.UUID) (model.Feed, error) {
	dbFeeds, err := feedService.Queries.MarkFeedFetched(ctx, feedId)
	if err != nil {
		return model.Feed{}, err
	}

	return model.DbFeedToFeed(&dbFeeds), nil
}
