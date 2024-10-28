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

type WSConnectionParams struct {
	ID          string
	Type        string
	CountryCode string
	IP          string
}

// Define the minimum compatible client version
const minimumClientVersion = "1.0.0"

// StartWebSocketServer starts a WebSocket server
func StartWebSocketServer(
	r *mux.Router,
	path string,
	messageChan chan []byte,
	ctx context.Context,
	connectCB func(params WSConnectionParams),
	disconnectCB func(params WSConnectionParams),
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

	// Fan-out messages to all connected clients
	go func() {
		for {
			select {
			case message := <-messageChan:
				log.Debug().
					Str("action", fmt.Sprintf("ws WRITE: %d", len(connections))).
					Str("payload", string(message)).
					Msgf("")
				func() {
					mutex.Lock()
					defer mutex.Unlock()
					for _, connWrapper := range connections {
						func() {
							connWrapper.mu.Lock()
							defer connWrapper.mu.Unlock()
							if err := connWrapper.conn.WriteMessage(websocket.TextMessage, message); err != nil {
								log.Error().Msgf("Error writing to websocket: %s", err.Error())
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

		// Read the initial message from the client, expected to be the version
		_, clientVersionMsg, err := conn.ReadMessage()
		if err != nil {
			log.Error().Msgf("Error reading client version: %s", err.Error())
			return
		}

		clientVersion := string(clientVersionMsg)
		log.Debug().Msgf("Client version received: %s", clientVersion)

		// Compare client version with minimum required version
		if clientVersion < minimumClientVersion {
			log.Warn().Msgf("Client version %s is incompatible, required: %s", clientVersion, minimumClientVersion)
			// Send a CloseMessage with code 1008 (Policy Violation) and a reason
			closeMessage := websocket.FormatCloseMessage(1008, clientVersion+" Please upgrade to the latest version: "+minimumClientVersion)
			conn.WriteMessage(websocket.CloseMessage, closeMessage)
			return // Exit to close the connection
		}

		conn.WriteMessage(websocket.TextMessage, []byte("version_accepted"))

		conn.SetPingHandler(nil)
		params := r.URL.Query()
		connParams := WSConnectionParams{
			ID:          params.Get("ID"),
			Type:        params.Get("Type"),
			CountryCode: r.Header.Get("Cf-Ipcountry"),
			IP:          r.Header.Get("Cf-Connecting-Ip"),
		}
		connectCB(connParams)
		addConnection(conn)

		log.Debug().
			Str("action", "⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪ ws CONNECT").
			Msgf("")
		for {
			messageType, _, err := conn.ReadMessage()
			if err != nil {
				log.Trace().Msgf("Client disconnected: %s", err.Error())
				removeConnection(conn)
				disconnectCB(connParams)
				break
			}
			if messageType == websocket.CloseMessage {
				log.Trace().Msgf("Received close frame from client.")
				removeConnection(conn)
				disconnectCB(connParams)
				break
			}
		}

		removeConnection(conn)
	})
}
