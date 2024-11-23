package http

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

const (
	initialDelay = 1.0
	maxAttempts  = 10
	exponential  = 1.2
)

// ConnectWebSocket establishes a new WebSocket connection
func ConnectWebSocket(url string, ctx context.Context) (chan []byte, error) {
	connectFactory := func() (*websocket.Conn, error) {
		currentBackoff := 0.0
		for attempt := 0; attempt < maxAttempts; attempt++ { // Add condition here
			select {
			case <-ctx.Done():
				return nil, fmt.Errorf("websocket connection canceled: %w", ctx.Err())
			default:
			}

			log.Debug().Msgf("WebSocket connection connecting: %s", url)
			conn, _, err := websocket.DefaultDialer.Dial(url, nil)
			if err != nil {
				var netErr *net.OpError
				if errors.As(err, &netErr) && netErr.Op == "dial" {
					log.Info().Msg("Solver service appears to be down")
					return nil, fmt.Errorf("solver service unavailable: %w", err)
				}

				log.Error().Msgf("WebSocket connection failed: %s\nReconnecting in %.3f seconds...", err, currentBackoff)
				timer := time.NewTimer(time.Duration(currentBackoff * float64(time.Second)))

				select {
				case <-ctx.Done():
					timer.Stop()
					return nil, fmt.Errorf("websocket connection canceled during backoff: %w", ctx.Err())
				case <-timer.C:
				}

				currentBackoff += initialDelay * math.Pow(exponential, float64(attempt))
				continue
			}

			conn.SetPongHandler(nil)
			return conn, nil
		}
		return nil, fmt.Errorf("maximum connection attempts (%d) reached", maxAttempts) // Add this return
	}

	pingInterval := time.NewTicker(time.Second * 5)
	connLk := &sync.Mutex{}
	responseCh := make(chan []byte)
	errCh := make(chan error)

	readMessage := func(conn *websocket.Conn) {
		defer close(responseCh)
		for {
			select {
			case <-ctx.Done():
				log.Info().Msg("Exiting readMessage loop due to context cancellation.")
				return
			default:
				messageType, p, err := conn.ReadMessage()
				if err != nil {
					errCh <- err
					return
				}
				if messageType == websocket.TextMessage {
					log.Debug().
						Str("action", "ws READ").
						Str("payload", string(p)).
						Msg("")
					responseCh <- p
				}
			}
		}
	}

	conn, err := connectFactory()
	if err != nil {
		log.Err(err).Msg("Error in WebSocket connection.")
		return nil, err
	}

	go readMessage(conn)

	go func() {
		defer pingInterval.Stop()
		for {
			select {
			case <-ctx.Done():
				log.Info().Msg("Ping loop exiting due to context cancellation.")
				return
			case <-pingInterval.C:
				connLk.Lock()
				log.Trace().Msg("Sending ping message.")
				if err := conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
					log.Err(err).Msg("Error sending ping message.")
					connLk.Unlock()
					return
				}
				connLk.Unlock()
			case err := <-errCh:
				log.Err(err).Msg("WebSocket error detected.")
				connLk.Lock()
				conn, err = connectFactory()
				connLk.Unlock()
				if err != nil {
					log.Err(err).Msg("Failed to reconnect WebSocket.")
					return
				}
				go readMessage(conn)
			}
		}
	}()

	return responseCh, nil
}
