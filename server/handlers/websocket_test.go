package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestWebSocketIntegration(t *testing.T) {
	l := log.New(os.Stdout, "test-ws", log.LstdFlags)
	wsHandler := NewWSHandler(l, nil)
	server := httptest.NewServer(http.HandlerFunc(wsHandler.HandleConnection))
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")

	t.Run("successful connection and message echo", func(t *testing.T) {
		ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
		if err != nil {
			t.Fatalf("could not open websocket connection: %v", err)
		}
		defer ws.Close()

		// test echo
		testMessage := "test message"
		err = ws.WriteMessage(websocket.TextMessage, []byte(testMessage))
		if err != nil {
			t.Fatalf("coudl not send message: %v", err)
		}

		// read response with timeout
		done := make(chan bool)
		go func() {
			_, msg, err := ws.ReadMessage()
			if err != nil {
				t.Errorf("could not read message %v", err)
				return
			}
			if string(msg) != testMessage {
				t.Errorf("expected message %s, got message %s", testMessage, string(msg))
			}
			done <- true
		}()

		select {
		case <-done:
		case <-time.After(2 * time.Second):
			t.Fatal("test timeout")
		}
	})

	t.Run("connection closes properly", func(t *testing.T) {
		ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
		if err != nil {
			t.Fatalf("could not open websocket connection %v", err)
		}

		if err := ws.Close(); err != nil {
			t.Errorf("error closing connection: %v", err)
		}

		err = ws.WriteMessage(websocket.TextMessage, []byte("test"))
		if err == nil {
			t.Error("expected error writing to closed ws connection")
		}
	})
}