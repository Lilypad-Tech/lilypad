package http

import (
	"context"
	"fmt"
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

type ConnectionWrapper struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

// StartWebSocketServer starts a WebSocket server
func StartWebSocketServer(
	r *mux.Router,
	path string,
	messageChan chan []byte,
	ctx context.Context,
) {
	var mutex = &sync.Mutex{}

	connections := map[*websocket.Conn]*ConnectionWrapper{}

	addConnection := func(conn *websocket.Conn) {
		mutex.Lock()
		defer mutex.Unlock()
		connections[conn] = &ConnectionWrapper{conn: conn}
	}

	removeConnection := func(conn *websocket.Conn) {
		mutex.Lock()
		defer mutex.Unlock()
		delete(connections, conn)
	}

	// spawn a reader from the incoming message channel
	// each message we get we fan out to all the currently connected websocket clients

	// TODO: we should add some subscription channels here because right now we are
	// splatting a lot of bytes down the write because everyone is hearing everything
	go func() {
		for {
			select {
			case message := <-messageChan:
				log.Debug().
					Str("action", fmt.Sprintf("ws WRITE: %d", len(connections))).
					Str("payload", string(message)).
					Msgf("")
				func() {
					// hold the mutex while we iterate over connections because
					// you can't modify a map while iterating over it (fatal
					// error: concurrent map iteration and map write)
					mutex.Lock()
					defer mutex.Unlock()
					for _, connWrapper := range connections {
						// wrap in a func so that we can defer the unlock so we can
						// unlock the mutex on panics as well as errors
						func() {
							connWrapper.mu.Lock()
							defer connWrapper.mu.Unlock()
							if err := connWrapper.conn.WriteMessage(websocket.TextMessage, message); err != nil {
								log.Error().Msgf("Error writing to websocket: %s", err.Error())
								// don't stop reading from messageChan just because one write failed
							}
						}()
					}
				}()
			case <-ctx.Done():
				return
			}
		}
	}()

	r.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Error().Msgf("Error upgrading websocket: %s", err.Error())
			return
		}
		defer conn.Close()
		addConnection(conn)

		log.Debug().
			Str("action", "⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪ ws CONNECT").
			Msgf("")
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

		removeConnection(conn)
	})
}
