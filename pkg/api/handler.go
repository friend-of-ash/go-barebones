package api

import (
	"net/http"

	"github.com/go-chi/render"
)

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["status"] = "Alive and kickin'"
	render.JSON(w, r, response)
}
