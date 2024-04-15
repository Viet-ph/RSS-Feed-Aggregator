package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/auth"
	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/service"
	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/utils"
)

type ContextKey string

const ContextUserKey ContextKey = "user"

func NewMiddlewareAuth(userService *service.UserService) func(http.Handler) http.Handler {
	//This will take the dependencies and return a authentication middleware that accepts only a single handler.
	//By doing this, will clean up the middleware function arguments and create the closure to outer deps.
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey, err := auth.GetAPIKey(r.Header)
			if err != nil {
				utils.RespondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Auth error: %v", err))
				return
			}

			user, err := userService.GetUserByAPIKey(r.Context(), apiKey)
			if err != nil {
				utils.RespondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Couldn't get user: %v", err))
				return
			}

			ctx := context.WithValue(r.Context(), ContextUserKey, user)
			handler.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
