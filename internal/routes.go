package internal

import (
	"net/http"

	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/handler"
	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/middleware"
	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/service"
)

func addRoutes(mux *http.ServeMux, userService *service.UserService) {
	middlewareAuth := middleware.NewMiddlewareAuth(userService)
	mux.HandleFunc("GET /api/healthz", handler.Readiness)
	mux.Handle("POST /v1/users", handler.HandleCreateUser(userService))
	mux.Handle("GET /v1/users", middlewareAuth(handler.HandleGetUserByAPIKey()))
}
