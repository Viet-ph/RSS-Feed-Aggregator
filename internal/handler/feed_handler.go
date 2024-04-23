package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/middleware"
	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/model"
	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/service"
	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/utils"
)

func HandleCreateFeed(feedService *service.FeedService, feedFollowService *service.FeedFollowService) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			type request struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			}

			req, err := utils.Decode[request](r)
			if err != nil {
				log.Printf("Error decoding parameters: %s", err)
				w.WriteHeader(500)
				return
			}

			user := r.Context().Value(middleware.ContextUserKey).(model.User) //Type assertion

			newFeed, err := feedService.CreateFeed(r.Context(), req.Name, req.Url, user.ID)
			if err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, "Error creating new feed")
				return
			}
			//Automatically follow feed when create new feed
			newFeedFollow, err := feedFollowService.CreateFeedFollow(r.Context(), user.ID, newFeed.ID)
			if err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating new feed follow: %v", err))
				return
			}

			utils.RespondWithJSON(w, http.StatusCreated, struct {
				Feed       model.Feed       `json:"feed"`
				FeedFollow model.FeedFollow `json:"feed_follow"`
			}{Feed: newFeed, FeedFollow: newFeedFollow})
		},
	)
}

func HandleGetFeeds(feedService *service.FeedService) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			feeds, err := feedService.GetFeeds(r.Context())
			if err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, "Error getting feeds")
				return
			}

			utils.RespondWithJSON(w, http.StatusOK, feeds)
		},
	)
}
