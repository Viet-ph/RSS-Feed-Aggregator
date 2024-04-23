package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/database"
	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/model"
	"github.com/google/uuid"
)

type PostService struct {
	Queries *database.Queries
}

func NewPostService(q *database.Queries) *PostService {
	return &PostService{
		Queries: q,
	}
}

func (postService *PostService) CreatePost(ctx context.Context, title, url, desc, pubDate string, feedId uuid.UUID) (model.Post, error) {
	description := sql.NullString{}
	if desc != "" {
		description.String = desc
		description.Valid = true
	}
	publishedAt := sql.NullTime{}
	if t, err := time.Parse(time.RFC1123Z, pubDate); err == nil {
		publishedAt = sql.NullTime{
			Time:  t,
			Valid: true,
		}
	}
	post, err := postService.Queries.CreatePost(ctx, database.CreatePostParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Title:       title,
		Url:         url,
		Description: description,
		PublishedAt: publishedAt,
		FeedID:      feedId,
	})

	if err != nil {
		return model.Post{}, err
	}

	return model.DbPostToPost(&post), nil
}

func (postService *PostService) GetPostsForUser(ctx context.Context, userId uuid.UUID, limit int32) ([]model.Post, error) {
	dbPosts, err := postService.Queries.GetPostsForUser(ctx, database.GetPostsForUserParams{
		UserID: userId,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}
	posts := make([]model.Post, 0, len(dbPosts))
	for _, dbPost := range dbPosts {
		posts = append(posts, model.DbPostToPost(&dbPost))
	}

	return posts, nil
}
