package http

import (
	"context"
	"time"

	"github.com/gorilla/websocket"
	"github.com/recws-org/recws"
	"github.com/rs/zerolog/log"
)

// ConnectWebSocket establishes a new WebSocket connection
func ConnectWebSocket(
	url string,
	messageChan chan []byte,
	ctx context.Context,
) *recws.RecConn {

	ws := &recws.RecConn{
		KeepAliveTimeout: 10 * time.Second,
	}
	ws.Dial(url, nil)

	go func() {
		for {
			select {
			case <-ctx.Done():
				go ws.Close()
				log.Printf("Websocket closed %s", ws.GetURL())
				return
			default:
				if !ws.IsConnected() {
					log.Printf("Websocket disconnected %s", ws.GetURL())
					continue
				}
				messageType, p, err := ws.ReadMessage()
				if err != nil {
					log.Printf("Error: ReadMessage %s", ws.GetURL())
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
		}
	}()
	return ws
}
