package main

import (
	"log"
	"net/http"
	"github.com/IwatsukaYura/speee_api/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)

	log.Println("Listing for requests at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
