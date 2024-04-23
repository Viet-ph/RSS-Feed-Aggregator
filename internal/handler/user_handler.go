package handler

import (
	"log"
	"net/http"

	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/middleware"
	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/service"
	"github.com/Viet-ph/RSS-Feed-Aggregator/internal/utils"
)

func HandleCreateUser(userService *service.UserService) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			type request struct {
				Name string `json:"name"`
			}

			req, err := utils.Decode[request](r)
			if err != nil {
				log.Printf("Error decoding parameters: %s", err)
				w.WriteHeader(500)
				return
			}

			newUser, err := userService.CreateUser(r.Context(), req.Name)
			if err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, "Error creating new user")
				return
			}

			utils.RespondWithJSON(w, http.StatusCreated, newUser)
		},
	)
}

func HandleGetUserByAPIKey() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			user := r.Context().Value(middleware.ContextUserKey)
			utils.RespondWithJSON(w, http.StatusOK, user)
		},
	)
}
