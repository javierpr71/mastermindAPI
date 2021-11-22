package controllers

import (
	"net/http"

	"github.com/javierpr71/mastermind/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "This is MasterMind Game API....")
}
