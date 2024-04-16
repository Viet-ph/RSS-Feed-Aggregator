package service

import (
	"context"
	"time"

	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/database"
	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/model"
	"github.com/google/uuid"
)

type FeedFollowService struct {
	Queries *database.Queries
}

func NewFeedFollowService(q *database.Queries) *FeedFollowService {
	return &FeedFollowService{
		Queries: q,
	}
}

func (feedFollowService *FeedFollowService) CreateFeedFollow(ctx context.Context, userId, feedId uuid.UUID) (*model.FeedFollow, error) {
	feedFollow, err := feedFollowService.Queries.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    userId,
		FeedID:    feedId,
	})

	if err != nil {
		return nil, err
	}

	return model.DbFeedFollowToFeedFollow(&feedFollow), nil
}

func (feedService *FeedFollowService) GetFeedFollows(ctx context.Context, userId uuid.UUID) ([]model.FeedFollow, error) {
	dbFeedFollows, err := feedService.Queries.GetFeedFollowsByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	feedFollows := make([]model.FeedFollow, 0, len(dbFeedFollows))
	for _, dbFeedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, *model.DbFeedFollowToFeedFollow(&dbFeedFollow))
	}

	return feedFollows, nil
}

func (feedService *FeedFollowService) DeleteFeedFollow(ctx context.Context, feedFollowId, userId uuid.UUID) error {
	err := feedService.Queries.DeleteFeedFollow(ctx, database.DeleteFeedFollowParams{
		ID:     feedFollowId,
		UserID: userId,
	})

	return err
}
