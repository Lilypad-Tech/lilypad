package http

import (
	"context"
	"net/http"
	"sync"

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
	var mutex = &sync.Mutex{}

	connections := map[*websocket.Conn]bool{}

	addConnection := func(conn *websocket.Conn) {
		mutex.Lock()
		defer mutex.Unlock()
		connections[conn] = true
	}

	removeConnection := func(conn *websocket.Conn) {
		mutex.Lock()
		defer mutex.Unlock()
		delete(connections, conn)
	}

	r.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Error().Msgf("Error upgrading websocket: %s", err.Error())
			return
		}
		defer conn.Close()
		addConnection(conn)

		wrappedCtx, wrappedCancel := context.WithCancel(ctx)

		go func() {
			for {
				select {
				case message := <-messageChan:
					log.Trace().
						Str("action", "ws WRITE").
						Str("payload", string(message)).
						Msgf("")
					for conn := range connections {
						c := conn
						go func() {
							if err := c.WriteMessage(websocket.TextMessage, message); err != nil {
								log.Error().Msgf("Error writing to websocket: %s", err.Error())
								return
							}
						}()
					}
				case <-wrappedCtx.Done():
					removeConnection(conn)
					return
				}
			}
		}()

		for {
			messageType, _, err := conn.ReadMessage()
			if err != nil {
				log.Trace().Msgf("Client disconnected: %s", err.Error())
				break
			}
			if messageType == websocket.CloseMessage {
				log.Trace().Msgf("Received close frame from client.")
				break
			}
		}

		wrappedCancel()

	})
}
