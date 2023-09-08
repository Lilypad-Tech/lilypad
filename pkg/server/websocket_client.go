package server

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// ConnectWebSocket establishes a new WebSocket connection
func ConnectWebSocket(url string, messageChan chan []byte) *websocket.Conn {
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

	// Read loop
	go func() {
		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				conn = ConnectWebSocket(url, messageChan)
				continue
			}
			if messageType == websocket.TextMessage {
				messageChan <- p
			}
		}
	}()

	return conn
}
