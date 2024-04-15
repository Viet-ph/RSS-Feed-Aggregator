package internal

import (
	"net/http"

	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/middleware"
	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/service"
)

func NewServer(userService *service.UserService) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, userService)
	var handler http.Handler = mux
	handler = middleware.MiddlewareCors(handler)
	return handler
}
