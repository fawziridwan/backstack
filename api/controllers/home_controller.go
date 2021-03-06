package controllers

import (
	"net/http"

	"github.com/fawziridwan/backstack/api/helpers"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	helpers.JsonResponse(w, http.StatusOK, "Welcome To This Awesome API")
}
