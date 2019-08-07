package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Dadinel/status-http/port"
	"github.com/Dadinel/status-http/status"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("", getRound).Methods("GET")
	router.HandleFunc("/{code}", getCode).Methods("GET")
	http.ListenAndServe(":"+port.GetPort(), router)
}

func getRound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(status.GetRandStatus())
}

func getCode(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(status.GetURLStatus(r))
}