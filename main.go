package main

import (
	"fmt"
	"net/http"

	"github.com/Dadinel/status-http/port"
	"github.com/Dadinel/status-http/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.SetRoutes(router)
	serverPort := port.GetPort()
	fmt.Printf(serverPort)
	http.ListenAndServe(":"+port.GetPort(), router)
}
