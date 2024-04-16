package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/middleware"
	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/model"
	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/service"
	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/utils"
)

func HandleCreateFeedFollow(feedFollowService *service.FeedFollowService) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			type request struct {
				FeedId uuid.UUID `json:"feed_id"`
			}

			req, err := utils.Decode[request](r)
			if err != nil {
				log.Printf("Error decoding parameters: %s", err)
				w.WriteHeader(500)
				return
			}
			user := r.Context().Value(middleware.ContextUserKey).(*model.User) //Type assertion

			newFeedFollow, err := feedFollowService.CreateFeedFollow(r.Context(), user.ID, req.FeedId)
			if err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating new feed follow: %v", err))
				return
			}

			utils.RespondWithJSON(w, http.StatusCreated, *newFeedFollow)
		},
	)
}

func HandleGetFeedFollows(feedFollowService *service.FeedFollowService) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			user := r.Context().Value(middleware.ContextUserKey).(*model.User) //Type assertion
			feedFollows, err := feedFollowService.GetFeedFollows(r.Context(), user.ID)
			if err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, "Error getting feed follows")
				return
			}

			utils.RespondWithJSON(w, http.StatusOK, feedFollows)
		},
	)
}

func HandleDeleteFeedFollow(feedFollowService *service.FeedFollowService) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			user := r.Context().Value(middleware.ContextUserKey).(*model.User) //Type assertion
			feedId, err := uuid.Parse(r.PathValue("feedFollowID"))
			if err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't parse feed follow id: %s", err))
				return
			}
			err = feedFollowService.DeleteFeedFollow(r.Context(), feedId, user.ID)
			if err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't delete feed follow: %s", err))
				return
			}

			utils.RespondWithJSON(w, http.StatusOK, struct{}{})
		},
	)
}
