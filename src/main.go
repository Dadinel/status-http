package main

import (
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("", getRound).Methods("GET")
	router.HandleFunc("/{code}", getCode).Methods("GET")
	//http.ListenAndServe(":"+getPort(), router)
}

// func getRound(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(getRandCode())
// }

// func getCode(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(getURLCode(r))
// }
