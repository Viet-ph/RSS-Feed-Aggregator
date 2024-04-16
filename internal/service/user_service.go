package service

import (
	"context"
	"time"

	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/database"
	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/model"
	"github.com/google/uuid"
)

type UserService struct {
	Queries *database.Queries
}

func NewUserService(q *database.Queries) *UserService {
	return &UserService{
		Queries: q,
	}
}

func (userService *UserService) CreateUser(ctx context.Context, username string) (*model.User, error) {
	user, err := userService.Queries.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      username,
	})

	if err != nil {
		return nil, err
	}

	return model.DbUserToUser(&user), nil
}

func (userService *UserService) GetUserByAPIKey(ctx context.Context, apiKey string) (*model.User, error) {
	user, err := userService.Queries.GetUserByAPIKey(ctx, apiKey)
	if err != nil {
		return &model.User{}, err
	}

	return model.DbUserToUser(&user), nil
}
