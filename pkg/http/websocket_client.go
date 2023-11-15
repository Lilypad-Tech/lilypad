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

	// if we ever get a cancellation from the context, try to close the connection
	go func() {
		<-ctx.Done()
		closed = true
		if conn != nil {
			conn.Close()
		}
	}()

	// retry connecting until we get a connection
	for {
		var err error
		log.Debug().Msgf("WebSocket connection connecting: %s", url)
		conn, _, err = websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			log.Error().Msgf("WebSocket connection failed: %s\nReconnecting in 2 seconds...", err)
			if closed {
				break
			}
			time.Sleep(2 * time.Second)
			continue
		}
		break
	}

	// now that we have a connection, if we haven't been closed yet, forever
	// read from the connection and send messages down the channel, unless we
	// fail a read in which case we try to reconnect
	if !closed {
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
					// exit this goroutine now, another one will be spawned if
					// the recursive call to ConnectWebSocket succeeds. Not
					// exiting this goroutine here will cause goroutines to pile
					// up forever concurrently calling conn.ReadMessage(), which
					// is not thread-safe.
					return
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
	}

	return conn
}
