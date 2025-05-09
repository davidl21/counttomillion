package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// initialize db
	serveMux := mux.NewRouter()

	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", serveMux)
	if err != nil {
		log.Fatal(err)
	}
}