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

type APIResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
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

func (c *Count) IncrementCount(rw http.ResponseWriter, r *http.Request) {
	newCount, err := c.store.IncrementCount(r.Context())
	if err != nil {
		c.l.Printf("Error incrementing count: %v", err)
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(APIResponse{
			Success: false,
			Message: "Error incrementing count in database",
			Data: nil,
		})
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(APIResponse{
		Success: true,
		Message: "Successfully incremented count",
		Data: newCount,
	})
}