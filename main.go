package main

import (
	"net/http"

	"github.com/Dadinel/status-http/port"
	"github.com/Dadinel/status-http/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.SetRoutes(router)
	http.ListenAndServe(":"+port.GetPort(), router)
}
