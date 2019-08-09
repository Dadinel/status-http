package main

import (
	"fmt"
	"net/http"

	"github.com/Dadinel/status-http/port"
	"github.com/Dadinel/status-http/routes"
	"github.com/gorilla/mux"
)

func main() {
	serverPort, err := port.GetPort()

	if err == nil {
		router := mux.NewRouter()
		routes.SetRoutes(router)
		http.ListenAndServe(":"+port.GetPort(), router)
	}
	else {
		fmt.Print(err)
	}
}
