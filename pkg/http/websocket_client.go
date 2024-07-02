package http

import (
	"context"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

// ConnectWebSocket establishes a new WebSocket connection
func ConnectWebSocket(
	url string,
	ctx context.Context,
) chan []byte {
	connectFactory := func() *websocket.Conn {
		for {
			log.Debug().Msgf("WebSocket connection connecting: %s", url)
			conn, _, err := websocket.DefaultDialer.Dial(url, nil)
			if err != nil {
				log.Error().Msgf("WebSocket connection failed: %s\nReconnecting in 2 seconds...", err)
				time.Sleep(2 * time.Second)
				continue
			}
			conn.SetPongHandler(nil)
			return conn
		}
	}

	pingInterval := time.NewTicker(time.Second * 5)
	connLk := &sync.Mutex{}
	responseCh := make(chan []byte)
	errCh := make(chan error)

	readMessage := func(conn *websocket.Conn) {
		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				errCh <- err
				return
			}
			if messageType == websocket.TextMessage {
				log.Debug().
					Str("action", "ws READ").
					Str("payload", string(p)).
					Msgf("")
				responseCh <- p
			}
		}
	}

	conn := connectFactory()
	go readMessage(conn)
	go func() {
		for {
			select {
			case <-pingInterval.C:
				connLk.Lock()
				log.Trace().Msg("send ping message")
				if err := conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
					log.Err(err).Msg("sending ping message")
					connLk.Unlock()
					continue
				}
				connLk.Unlock()
			case err := <-errCh:
				log.Err(err).Msg("websocket error")
				connLk.Lock()
				conn = connectFactory()
				connLk.Unlock()
				go readMessage(conn)
			}
		}
	}()
	return responseCh
}
