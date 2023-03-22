package app

import (
	"encoding/json"
	"net/http"
)

type HelloWorld struct {
	Message string `json:"message"`
}

func (a *App) CreateTeamHandler(w http.ResponseWriter, r *http.Request) {
	payload := &HelloWorld{
		Message: "Hello World",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode((payload))
}
