package http

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// StartWebSocketServer starts a WebSocket server
func StartWebSocketServer(
	r *mux.Router,
	path string,
	messageChan chan []byte,
	ctx context.Context,
) {
	connections := map[*websocket.Conn]bool{}

	r.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Error().Msgf("Error upgrading websocket: %s", err.Error())
			return
		}
		defer conn.Close()
		connections[conn] = true

		wrappedCtx, wrappedCancel := context.WithCancel(ctx)

		go func() {
			for {
				select {
				case message := <-messageChan:
					for conn := range connections {
						if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
							log.Error().Msgf("Error writing to websocket: %s", err.Error())
							return
						}
					}
				case <-wrappedCtx.Done():
					delete(connections, conn)
					return
				}
			}
		}()

		for {
			messageType, _, err := conn.ReadMessage()
			if err != nil {
				log.Info().Msgf("Client disconnected: %s", err.Error())
				break
			}
			if messageType == websocket.CloseMessage {
				log.Info().Msgf("Received close frame from client.")
				break
			}
		}

		wrappedCancel()

	})
}
