package http

import (
	"context"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// ConnectWebSocket establishes a new WebSocket connection
func ConnectWebSocket(
	url string,
	messageChan chan []byte,
	ctx context.Context,
) *websocket.Conn {
	closed := false
	var conn *websocket.Conn
	for {
		var err error
		conn, _, err = websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			log.Println("WebSocket connection failed:", err, "Reconnecting in 2 seconds...")
			time.Sleep(2 * time.Second)
			continue
		}
		break
	}
	log.Println("WebSocket connected")

	conn.Close()
	// Read loop
	go func() {
		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				if closed {
					return
				}
				log.Println("Read error:", err)
				conn = ConnectWebSocket(url, messageChan, ctx)
				continue
			}
			if messageType == websocket.TextMessage {
				messageChan <- p
			}
		}
	}()

	// Wait for close
	go func() {
		for {
			select {
			case <-ctx.Done():
				closed = true
				return
			}
		}
	}()

	return conn
}
