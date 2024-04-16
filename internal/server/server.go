package server

import (
	"net/http"

	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/middleware"
	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/service"
)

func NewServer(
	userService *service.UserService,
	feedService *service.FeedService,
	feedFollowService *service.FeedFollowService,
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux,
		userService,
		feedService,
		feedFollowService,
	)
	var handler http.Handler = mux
	handler = middleware.MiddlewareCors(handler)
	return handler
}
