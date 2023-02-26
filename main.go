package main

import (
	"log"
	"net/http"

	"github.com/felipeas314/go-api-with-mux/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//configs.ConnectDB()

	routes.UserRoute(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
