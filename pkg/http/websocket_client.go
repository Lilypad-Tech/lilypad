package http

import (
	"context"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
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
		log.Info().Msgf("WebSocket connection connecting: %s", url)
		conn, _, err = websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			log.Error().Msgf("WebSocket connection failed: %s\nReconnecting in 2 seconds...", err)
			time.Sleep(2 * time.Second)
			continue
		}
		break
	}
	log.Info().Msgf("WebSocket connected")

	go func() {
		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				if closed {
					return
				}
				log.Error().Msgf("Read error: %s\nReconnecting in 2 seconds...", err)
				time.Sleep(2 * time.Second)
				conn = ConnectWebSocket(url, messageChan, ctx)
				continue
			}
			if messageType == websocket.TextMessage {
				log.Debug().
					Str("action", "ws READ").
					Str("payload", string(p)).
					Msgf("")
				messageChan <- p
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				closed = true
				conn.Close()
				return
			}
		}
	}()

	return conn
}
