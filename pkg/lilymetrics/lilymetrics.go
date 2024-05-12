package lilymetrics

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	// "net/http"

	// "github.com/elazarl/goproxy/transport"
	"github.com/gorilla/mux"

	// "github.com/gorilla/websocket"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"

	// "go.opentelemetry.io/otel"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

type Event struct {
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	Details   string `json:"details"`
}
type Update struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

var db *sql.DB
var logger zerolog.Logger
var sqlLogFile os.File
var server *socketio.Server

// handleEvent listens for incoming event data and logs it into the PostgreSQL database.
func generateLogFileName() string {
	currentTime := time.Now()
	return "migrations/tmp/lilylog_" + currentTime.Format("2006-01-02_15") + ".log"
}
func generateSqlFileName() string {
	currentTime := time.Now()
	return "migrations/tmp/" + currentTime.Format("200601021504") + "_logs.up.sql"
}

func (e Event) String() string {
	// convert the Event to a string
	// for example, you could marshal it to JSON:
	bytes, err := json.Marshal(e)
	if err != nil {
		return "error marshaling Event to JSON: " + err.Error()
	}
	return string(bytes)
}
func handleJobEvent(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var event Event
	// var event map[string]interface{}
	err := decoder.Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	logger.Info().Msg(event.String())
	fmt.Println("Received event", event)
}

func handleEvent(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var event Event
	// var event map[string]interface{}
	err := decoder.Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	event.Timestamp = time.Now().UTC().Format(time.RFC3339Nano)
	logger.Info().Msg(event.Details)

	log.Printf("Event Received and Logged: %+v", event)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Event received and logged successfully"))
}

type Log struct {
	id        string
	timestamp string
	details   string
}

func handleGetEvent(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte("{result:'ok'}"))
	// Create a Response struct
	// updates := []Update{
	// 	{ID: "1", Message: "Update 1"},
	// 	{ID: "2", Message: "Update 2"},
	// 	{ID: "3", Message: "Update 3"},
	// }
	// Handle error
	// Close the rows when the function exits
	// Create a slice to hold updates
	// Iterate through the rows
	// Create a variable to hold each log
	// Scan the values from the row into the Log struct fields
	// Handle error
	// Convert the Log into an Update and append it to the updates slice
	// Check for any errors during iteration
	// Handle error
	updates := getLatestLogs()
	// Create a map to hold the JSON data
	// responseMap := map[string]interface{}{
	// 	"updates": updates,
	// }

	// Marshal the map into JSON
	jsonData, err := json.Marshal(updates)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON data to the response writer
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func getLatestLogs() []Update {
	rows, err := db.Query("SELECT  id, timestamp,details FROM logs ORDER BY timestamp limit 10")
	if err != nil {

	}

	defer rows.Close()

	var updates []Update

	for rows.Next() {

		var logs Log

		if err := rows.Scan(&logs.id, &logs.timestamp, &logs.details); err != nil {

			log.Println("Error scanning row: %v", err)
		}

		updates = append(updates, Update{ID: logs.id, Message: logs.details})
	}

	if err := rows.Err(); err != nil {

	}
	return updates
}
func escapeSingleQuotes(s string) string {
	return fmt.Sprintf("%s", s)
}

// func newConsoleExporter() (oteltrace.SpanExporter, error) {
// 	return stdouttrace.New()
// 	// return nil, nil
// }
// func newTraceProvider(exp sdktrace.SpanExporter) *sdktrace.TracerProvider {
// 	// Ensure default SDK resources and the required service name are set.
// 	r, err := resource.Merge(
// 		resource.Default(),
// 		resource.NewWithAttributes(
// 			semconv.SchemaURL,
// 			semconv.ServiceName("Metrics"),
// 		),
// 	)

// 	if err != nil {
// 		panic(err)
// 	}

//		return sdktrace.NewTracerProvider(
//			sdktrace.WithBatcher(exp),
//			sdktrace.WithResource(r),
//		)
//	}
func initializeListener() *pq.Listener {
	// connStr := "postgres://user:password@localhost/database_name?sslmode=disable"
	dbHost := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := "postgres" //os.Getenv("POSTGRES_DB")
	connStr := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"

	listener := pq.NewListener(connStr, 10*time.Second, time.Minute, func(ev pq.ListenerEventType, err error) {
		if err != nil {
			log.Fatalln(err)
			// log.Error().Err(err).Msg("PostgreSQL listener error")
		}
	})

	err := listener.Listen("updates")
	if err != nil {
		log.Fatalln(err)
		// log.Fatal().Err(err).Msg("Error listening for PostgreSQL notifications")
	}
	log.Println("RETURN Listener")
	return listener
}

func handleNotifications(server *socketio.Server, l *pq.Listener) {
	for {
		log.Println("update")
		select {
		case <-l.Notify:
			fmt.Println("received notification, new work available")
			// updates := []Update{
			// 	{ID: "1", Message: "Update 1"},
			// 	{ID: "2", Message: "Update 2"},
			// 	{ID: "3", Message: "Update 3"},
			// }
			updates := getLatestLogs()
			emitSocketEvent("updates", updates)
		case <-time.After(90 * time.Second):
			go l.Ping()
			// Check if there's more work available, just in case it takes
			// a while for the Listener to notice connection loss and
			// reconnect.
			fmt.Println("received no work for 90 seconds, checking for new work")
		}
		// select {
		// case notification := <-listener.Notify:
		// 	// Handle PostgreSQL notification
		// 	//log.Info().Interface("Notification", notification).Msg("Received notification")
		// 	log.Println("update")
		// 	_ = notification
		// 	// Emit a socket event with the received notification data
		// 	// emitSocketEvent("notification", notification)
		// }
	}
}

// var server *socketio.Server

func emitSocketEvent(eventName string, data []Update) { // data *pq.Notification
	// Emit the event to all connected clients
	server.BroadcastToNamespace("/", "update", data, "data")
	// if err != nil {
	// 	log.Println("Error broadcasting event to clients")
	// }
}
func RunMetrics() {
	// defer Trace(context.Background(), "Migrations").End()
	span := TraceSection(context.Background(), "Migrations")
	MigrateUp("logs_bulk")
	MigrateUp("schema")
	span.End()
	dbHost := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := "postgres" //os.Getenv("POSTGRES_DB")
	connStr := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		//log.Fatal(err)
		fmt.Println("Error opening database connection")
	}

	// Generate the initial log file name based on the current time
	logFileName := generateLogFileName()
	sqlLogFileName := generateSqlFileName()
	// Open the initial log file
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// sqlLogFile, err := os.OpenFile(sqlLogFileName, os.O_CREATE|os.O_APPEND, 0666)
	sqlLogFile, err := os.Create(sqlLogFileName)
	if err != nil {
		fmt.Println("Failed to open log file")
	}

	// Initialize the logger with the initial log file
	logger = zerolog.New(logFile).With().Timestamp().Logger().Hook(zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, msg string) {
		currentTime := time.Now().UTC().Format(time.RFC3339Nano)
		_, err = db.Exec("INSERT INTO logs (type, timestamp, details) VALUES ($1, $2, $3)", level, currentTime, msg)
		if err != nil {
			log.Printf("Error inserting event into database: %v", err)
		}
		sql := fmt.Sprintf(`INSERT INTO logs_bulk (type, timestamp, details) VALUES ('%s', '%s', '%s');`, level, currentTime, strings.ReplaceAll(msg, "'", "\""))
		sqlLogFile.WriteString(sql + "\n")
		fmt.Println("should log " + sql)

	}))

	// Start a goroutine to periodically check and update the log file name
	go func() {
		for {
			// Sleep until the next hour
			nextHour := time.Now().Truncate(time.Minute).Add(time.Minute)
			time.Sleep(time.Until(nextHour))

			// Generate the new log file name
			newLogFileName := generateLogFileName()
			newSqlLogFileName := generateSqlFileName()
			// Close the current log file
			if err := logFile.Close(); err != nil {
				fmt.Println("Failed to close log file")
			}
			fileInfo, err := sqlLogFile.Stat()
			if err != nil {
				fmt.Println("Error getting file information:", err)
				return
			}
			size := fileInfo.Size()
			if err := sqlLogFile.Close(); err != nil {
				fmt.Println("Failed to close sql log file")
			}

			// Get file information
			if size == 0 {
				os.Remove(sqlLogFile.Name()) // todo: instead remove all empty files
			} else {
				os.Rename(sqlLogFile.Name(), "migrations/logs_bulk/"+filepath.Base(sqlLogFile.Name()))
			}

			// Open the new log file
			newLogFile, err := os.OpenFile(newLogFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				log.Println("Failed to open new log file")
			}
			newSqlLogFile, err := os.Create(newSqlLogFileName) //os.OpenFile(newSqlLogFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				log.Println("Failed to open new log file")
			}

			// Update the logger to use the new log file
			// Assign the new log file to the logFile variable
			logFile = newLogFile
			sqlLogFile = newSqlLogFile
		}
	}()

	stupRoutes()

}

var allowOriginFunc = func(r *http.Request) bool {
	return true
}

func stupRoutes() {
	server = socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})
	listener := initializeListener()
	go handleNotifications(server, listener)
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		log.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		log.Println("chat:", msg)
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()

	router := mux.NewRouter()
	// http.Handle("/socket.io/", server)
	router.HandleFunc("/socket.io/", func(w http.ResponseWriter, r *http.Request) {
		server.ServeHTTP(w, r)
	})
	// router.Handle("/socket.io/", func(context echo.Context) error {
	// 	server.ServeHTTP(context.Response(), context.Request())
	// 	return nil
	// })
	router.HandleFunc("/metrics-dashboard/log", handleEvent).Methods("POST") // Only keep the event handler route
	router.HandleFunc("/metrics-dashboard/job", handleJobEvent).Methods("POST")
	// router.HandleFunc("/", http.FileServer(http.Dir("../frontend"))).Methods("GET") // Only keep the event handler route

	router.HandleFunc("/updates", handleGetEvent).Methods("GET")
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("dashboard/build"))))
	http.Handle("/", http.FileServer(http.Dir("dashboard/build")))
	log.Fatal(http.ListenAndServe(":8001", router))
}
func _Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", allowHeaders)

		next.ServeHTTP(w, r)
	})
}
func main() {
	RunMetrics()

}

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }
