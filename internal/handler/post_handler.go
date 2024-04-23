package handler

import (
	"net/http"
	"strconv"

	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/middleware"
	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/model"
	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/service"
	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/utils"
)

func HandleGetPostsByUser(postService *service.PostService) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			user := r.Context().Value(middleware.ContextUserKey).(model.User) //Type assertion

			limitStr := r.URL.Query().Get("limit")
			var limit int
			if limitStr == "" {
				limit = 5
			} else {
				var err error
				limit, err = strconv.Atoi(limitStr)
				if err != nil {
					utils.RespondWithError(w, http.StatusInternalServerError, "Error convert limit query param to int")
					return
				}
			}
			posts, err := postService.GetPostsForUser(r.Context(), user.ID, int32(limit))
			if err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, "Error getting posts")
				return
			}

			utils.RespondWithJSON(w, http.StatusOK, posts)
		},
	)
}
