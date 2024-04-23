package handler

import (
	"log"
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
			log.Println(limitStr)
			limit, err := strconv.Atoi(limitStr)
			if err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, "Error getting limit query param")
				return
			}
			if limit == 0 {
				limit = 5
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
