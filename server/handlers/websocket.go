package handlers

import (
	"log"
	"net/http"
	"sync"

	"github.com/davidl21/counttomillion/server/data"
	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	l *log.Logger
	store *data.Store
	clients map[*websocket.Conn]bool
	broadcast chan []byte
	mu *sync.Mutex
	upgrader websocket.Upgrader
}

func NewWSHandler(l *log.Logger, s *data.Store) *WebSocketHandler {
	return &WebSocketHandler{
		l: l,
		store: s,
		clients: make(map[*websocket.Conn]bool),
		broadcast: make(chan []byte),
		mu: &sync.Mutex{},
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (ws *WebSocketHandler) HandleConnection(rw http.ResponseWriter, r *http.Request) {
	conn, err := ws.upgrader.Upgrade(rw, r, nil)
	if err != nil {
		ws.l.Println("Error upgrading:", err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			ws.l.Println("Error reading message:", err)
			break
		}
		ws.l.Println("Message received", message)
		
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
		ws.l.Println("Error writing message:", err)
			break
		}
	}
}