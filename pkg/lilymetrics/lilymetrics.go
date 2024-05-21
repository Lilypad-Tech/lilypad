package lilymetrics

import (
	"bufio"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	// "net/http"

	// "github.com/elazarl/goproxy/transport"

	// "github.com/gorilla/websocket"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"

	// "go.opentelemetry.io/otel"
	socketio "github.com/googollee/go-socket.io"
)

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

func handleMatcherEvent(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	data := string(body)
	server.BroadcastToNamespace("/", "matcher", data, "data")
	fmt.Println("Received data:", data)
}
func handleJobResult(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	b := string(body)

	server.BroadcastToNamespace("/", "resultfolder", b, "data")
	return
	fmt.Println("Received data:", b)
	if strings.Contains(b, "Complete") {
		fmt.Println("status: Complete")
		type Data struct {
			Dealid   string `json:"dealid"`
			State    string `json:"state"`
			Metadata string `json:"metadata"`
		}

		var data Data
		err = json.Unmarshal([]byte(b), &data)
		if err != nil {
			fmt.Println("Error unmarshalling JSON data")
		}

		jobresult := data.Dealid
		cmdstd := exec.Command("cat", "/tmp/lilypad/data/downloaded-files/"+jobresult+"/stdout")

		output, err := cmdstd.Output()
		o := string(output)
		if err != nil {
			fmt.Println("Error:", err)
		}

		server.BroadcastToNamespace("/", "result", o, "data")

		fmt.Println("output", string(output))

	}
}

func handleJobEvent(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	b := string(body)

	server.BroadcastToNamespace("/", "deal", b, "data")
	fmt.Println("Received data:", b)
	// currentTime := time.Now().UTC().Format(time.RFC3339Nano)
	// _, err = db.Exec("INSERT INTO jobs (type, timestamp, details) VALUES ($1, $2, $3)", level, currentTime, msg)

	if err != nil {
		log.Printf("Error inserting event into database: %v", err)
	}
	if strings.Contains(b, "Complete") {
		fmt.Println("status: Complete")
		// r := data.dealid
		// server.BroadcastToNamespace("/", "result", r, "data")
		type Data struct {
			Dealid   string `json:"dealid"`
			State    string `json:"state"`
			Metadata string `json:"metadata"`
		}

		var data Data

		err = json.Unmarshal([]byte(b), &data)
		if err != nil {
			fmt.Println("Error unmarshalling JSON data")
			// http.Error(w, err.Error(), http.StatusBadRequest)
			// return
		}

		//server.BroadcastToNamespace("/", "result", string(body), "data")
		// fmt.Println("ResultsAccepted:", string(body))
		// if strings.Contains(string(body), "ResultsAccepted") {
		jobresult := data.Dealid
		cmdstd := exec.Command("cat", "/tmp/lilypad/data/downloaded-files/"+jobresult+"/stdout")

		output, err := cmdstd.Output()
		o := string(output)
		if err != nil {
			fmt.Println("Error:", err)
			//return true
		}
		server.BroadcastToNamespace("/", "result", o, "data")
		fmt.Println(string(output))
	}

}
func handleStatusEvent(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	b := string(body)

	server.BroadcastToNamespace("/", "deal", b, "data")
	fmt.Println("Received data:", b)
	var data JobUpdate

	err = json.Unmarshal([]byte(b), &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON data")

	}
	currentTime := time.Now().UTC().Format(time.RFC3339Nano)
	// Define your function to insert or update a job
	insertOrUpdateJob := func(jobID, status, module_id string) error {
		// Check if job_id exists
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM jobs WHERE job_id = $1", jobID).Scan(&count)
		if err != nil {
			return err
		}

		if count == 0 {
			// Insert new job if job_id does not exist
			_, err := db.Exec("INSERT INTO jobs (job_id, status, time_start, module_id,details) VALUES ($1, $2, $3, $4, $5)", jobID, status, currentTime, module_id, "{}")
			return err
		}

		// Update status and timeupdated if job_id exists
		_, err = db.Exec("UPDATE jobs SET status = $1, time_update = $2 WHERE job_id = $3", status, currentTime, jobID)
		return err
	}
	err = insertOrUpdateJob(data.JobID, data.Status, data.ModuleID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
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

func handleGetJobUpdates(w http.ResponseWriter, r *http.Request) {

	updates := getLatestJobs()

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
func handleGetLogUpdates(w http.ResponseWriter, r *http.Request) {

	updates := getLatestLogs()

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
func getLatestJobs() []JobUpdate {
	rows, err := db.Query("SELECT id, module_id, job_id, status, time_update, time_start, details FROM jobs ORDER BY  time_start desc limit 5")
	if err != nil {

	}

	defer rows.Close()

	var updates []JobUpdate

	for rows.Next() {

		var jobs DbJob

		if err := rows.Scan(&jobs.id, &jobs.module_id, &jobs.job_id, &jobs.status, &jobs.time_update, &jobs.time_start, &jobs.details); err != nil {

			log.Println("Error scanning row: %v", err)
		}

		updates = append(updates, JobUpdate{ID: jobs.id, ModuleID: jobs.module_id, JobID: jobs.job_id, Status: jobs.status, TimeUpdate: jobs.time_update, TimeStart: jobs.time_start, Details: jobs.details})
	}

	if err := rows.Err(); err != nil {

	}

	return updates
}
func getLatestLogs() []Update {
	rows, err := db.Query("SELECT  id, timestamp,details FROM logs ORDER BY  timestamp desc limit 5")
	if err != nil {

	}

	defer rows.Close()

	var updates []Update

	for rows.Next() {

		var logs Log

		if err := rows.Scan(&logs.id, &logs.timestamp, &logs.details); err != nil {

			log.Println("Error scanning row: %v", err)
		}

		updates = append(updates, Update{ID: logs.id, Timestamp: logs.timestamp, Message: logs.details})
	}

	if err := rows.Err(); err != nil {

	}

	return updates
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
	return listener
}

func handleNotifications(server *socketio.Server, l *pq.Listener) {
	for {
		// log.Println("update")
		select {
		case <-l.Notify:
			fmt.Println("received notification, new work available")

			updates := getLatestLogs()
			emitSocketEvent("updates", updates)
			// case <-time.After(90 * time.Second):
			// 	go l.Ping()
			// 	// Check if there's more work available, just in case it takes
			// 	// a while for the Listener to notice connection loss and
			// 	// reconnect.
			// 	fmt.Println("received no work for 90 seconds, checking for new work")
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
}
func RunMetrics() {
	span := TraceSection(context.Background(), "Migrations")
	MigrateUp("schema")
	MigrateUp("logs_bulk")
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

	setupRoutes()

}

func runTask(msg string) bool {
	log.Println("notice:", msg)
	server.BroadcastToNamespace("/", "reply", msg, "data")
	words := strings.Fields(msg)
	binaryName := os.Getenv("PWD") + "/lilypad"
	cobraCommand := "run"
	cobraArgs := []string{words[0], words[1], strings.Join(words[2:], " ")} // Example arguments

	// Combine the binary name, Cobra command, and arguments
	command := append([]string{binaryName, cobraCommand}, cobraArgs...)

	// Execute the Cobra command
	cmd := exec.Command(command[0], command[1:]...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		r := scanner.Text()
		fmt.Println("from scanner ", r)
	}

	if err := scanner.Err(); err != nil {
		//log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		//log.Fatal(err)
	}
	return false
}

func main() {
	RunMetrics()
}
