package routes

import (
	"net/http"

	"github.com/Dadinel/status-http/status"
	"github.com/gorilla/mux"
)

// SetRoutes : Define os endpoints
func SetRoutes(router *mux.Router) {
	router.HandleFunc("/", getRound)
	router.HandleFunc("/{code}", getCode)
}

func getRound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(status.GetRandStatus())
}

func getCode(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(status.GetURLStatus(r))
}
