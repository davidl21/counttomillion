package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/davidl21/counttomillion/server/data"
	"github.com/davidl21/counttomillion/server/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// initialize db
	ctx := context.Background()
	dbURL := os.Getenv("DATABASE_URL")

	store, err := data.NewStore(ctx, dbURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Ping(ctx); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully pinged database")
	}
	defer store.Close()

	l := log.New(os.Stdout, "count-api ", log.LstdFlags)

	countHandler := handlers.NewCount(l, store)
	wsHandler := handlers.NewWSHandler(l, store)

	serveMux := mux.NewRouter()
	serveMux.HandleFunc("/ws", wsHandler.HandleConnection)

	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/count", countHandler.IncrementCount)

	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/count", countHandler.IncrementCount)

	l.Println("Starting server on port 8080")
	err = http.ListenAndServe(":8080", serveMux)
	if err != nil {
		log.Fatal(err)
	}
}