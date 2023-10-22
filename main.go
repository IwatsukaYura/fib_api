package main

import (
	"log"
	"net/http"

	"github.com/IwatsukaYura/fib_api/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/fib", handlers.FibonacciHandler).Methods(http.MethodGet)

	log.Println("Listing for requests")
	log.Fatal(http.ListenAndServe(":8080", r))
}
