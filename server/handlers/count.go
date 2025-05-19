package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/davidl21/counttomillion/server/data"
)

type Count struct {
	l *log.Logger
	store *data.Store
}

func NewCount(l *log.Logger, store *data.Store) *Count {
	return &Count{l, store}
}

func (c *Count) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	c.l.Println("Connected to Count handler")

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]string{
		"message": "successfully connected to count handler",
	})
}