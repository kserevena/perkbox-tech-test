package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	storage := initStorage()

	router := mux.NewRouter()
	router = initRoutes(router, storage)
	log.Fatal(http.ListenAndServe(":8000", router))
}
