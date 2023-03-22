package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Tomdango/retrotool-api-lambda-v1/lib/auth"
	team "github.com/Tomdango/retrotool-api-lambda-v1/lib/db/teams"
	"github.com/google/uuid"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string    `json:"message"`
	Team    team.Team `json:"team"`
}

func (handler *Handler) CreateTeam(w http.ResponseWriter, r *http.Request) {
	token, err := auth.ParseJWTFromHeaders(r.Header, handler.Config.UserPoolID)
	if err != nil {
		payload := &ErrorResponse{
			Message: err.Error(),
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(payload)
		return
	}

	svc, err := team.NewTeamService("prod-teams-1")

	if err != nil {
		payload := &ErrorResponse{
			Message: err.Error(),
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(payload)
		return
	}

	newTeam := &team.Team{
		ID:      uuid.NewString(),
		Name:    "Test Team",
		Members: []team.Member{{ID: token.Subject()}},
	}

	err = svc.Create(newTeam)
	if err != nil {
		payload := &ErrorResponse{
			Message: err.Error(),
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(payload)
		return
	}

	payload := &SuccessResponse{
		Message: "Successfully Created Team",
		Team:    *newTeam,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payload)
}

func (handler *Handler) GetTeamByID(w http.ResponseWriter, r *http.Request) {
	token, err := auth.ParseJWTFromHeaders(r.Header, handler.Config.UserPoolID)

	
}
