package lilymetrics

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"

	_ "github.com/lib/pq"
)

type Event struct {
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	Details   string `json:"details"`
}

var db *sql.DB
var logger zerolog.Logger
var sqlLogFile os.File

// handleEvent listens for incoming event data and logs it into the PostgreSQL database.
func generateLogFileName() string {
	currentTime := time.Now()
	return "app_" + currentTime.Format("2006-01-02_15") + ".log"
}
func generateSqlFileName() string {
	currentTime := time.Now()
	return "migrations/tmp/" + currentTime.Format("200601021504") + "_logs.up.sql"
}

//	func generateSqlFileName() string {
//		currentTime := time.Now()
//		return "migrations/tmp/" + currentTime.Format("2006010215") + "_data.sql"
//	}
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
	// event.Details = "This is a test event"
	// event.Type = event.

	fmt.Println("Received event")
	// if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
	// 	http.Error(w, "Bad request", http.StatusBadRequest)
	// 	return
	// }

	// var event map[string]interface{}
	// err := decoder.Decode(&event)
	// if err != nil {
	//     http.Error(w, err.Error(), http.StatusBadRequest)
	//     return
	// }
	logger.Info().Msg(event.Details)
	// Now you can use the event map
	fmt.Println("details:", event.Details)
	// defer r.Body.Close()

	// r = fmt.Sprintf(`{"Type":"%s","Details":"%s"}`, event.Type, event.Details)
	// fmt.Println(r)

	// sql := fmt.Sprintf(`"INSERT INTO logs (type, timestamp, details) VALUES ("%s", "%s","%s")"`, event.Type, event.Timestamp, event.Details)
	// fmt.Println(sql)
	// _, err = db.Exec("INSERT INTO logs (type, timestamp, details) VALUES ($1, $2, $3)", event.Type, event.Timestamp, event.Details)

	if err != nil {
		log.Printf("Error inserting event into database: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("Event Received and Logged: %+v", event)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Event received and logged successfully"))
}
func handleGetEvent(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
func escapeSingleQuotes(s string) string {
	return fmt.Sprintf("%s", s)
}
func RunMetrics() {

	dbHost := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := "postgres" //os.Getenv("POSTGRES_DB")
	connStr := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"
	fmt.Println("connStr: ", connStr)
	MigrateUp("logs_bulk")
	MigrateUp("schema")

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/metrics-dashboard/log", handleEvent).Methods("POST") // Only keep the event handler route
	router.HandleFunc("/metrics-dashboard/job", handleJobEvent).Methods("POST")
	router.HandleFunc("/", handleGetEvent).Methods("GET") // Only keep the event handler route

	// Generate the initial log file name based on the current time
	logFileName := generateLogFileName()
	sqlLogFileName := generateSqlFileName()
	// Open the initial log file
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// sqlLogFile, err := os.OpenFile(sqlLogFileName, os.O_CREATE|os.O_APPEND, 0666)
	sqlLogFile, err := os.Create(sqlLogFileName)
	// sqlLogFile.WriteString("test\n")
	// sqlLogFile.Close()
	if err != nil {
		// log.Fatal().Err(err).Msg("Failed to open log file")
		// sql := fmt.Sprintf(`"INSERT INTO logs (type, timestamp, details) VALUES ("%s", "%s","%s")"`, event.Type, event.Timestamp, event.Details)
		// fmt.Println(sql)
		fmt.Println("Failed to open log file")
	}

	// Initialize the logger with the initial log file
	logger = zerolog.New(logFile).With().Timestamp().Logger().Hook(zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, msg string) {
		// fmt.Print(e)
		// contextData, ok := e.GetCtx().Value("timestamp").(string)
		// if ok {
		// 	e.Str("timestamp", contextData)
		// }
		currentTime := time.Now().UTC().Format(time.RFC3339Nano)

		// table_name := "logs"
		// sql := fmt.Sprintf(`INSERT INTO logs (type, timestamp, details) VALUES ('%s', '%s', '%s');`, level, currentTime, strings.ReplaceAll(msg, "'", "\""))
		_, err = db.Exec("INSERT INTO logs (type, timestamp, details) VALUES ($1, $2, $3)", level, currentTime, msg)
		if err != nil {
			log.Printf("Error inserting event into database: %v", err)
		}
		sql := fmt.Sprintf(`INSERT INTO logs_bulk (type, timestamp, details) VALUES ('%s', '%s', '%s');`, level, currentTime, strings.ReplaceAll(msg, "'", "\""))
		sqlLogFile.WriteString(sql + "\n")
		//  := e.Str("timestamp")
		// event := Event{
		// 	Type:      "your type here",
		// 	Timestamp: "your timestamp here",
		// 	Details:   msg,
		// }

		// eventJson, _ := json.Marshal(event)
		// timestamp, _ := time.Parse(time.RFC3339, timestampStr)
		//sql := fmt.Sprintf(`"INSERT INTO logs (type, timestamp, details) VALUES ("%s", "%s","%s")"`, e.Type, e.Timestamp, "e.Details")

		// _, err = db.Exec("INSERT INTO logs (type, timestamp, details) VALUES ($1, $2, $3)", event.Type, event.Timestamp, event.Details)

		fmt.Println("should log " + sql)

	}))
	// logger.Info().Msg("test")
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
				os.Remove(sqlLogFile.Name())
			} else {
				os.Rename(sqlLogFile.Name(), "migrations/logs_bulk/"+filepath.Base(sqlLogFile.Name()))
			}

			// Open the new log file
			newLogFile, err := os.OpenFile(newLogFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				fmt.Println("Failed to open new log file")
			}
			newSqlLogFile, err := os.Create(newSqlLogFileName) //os.OpenFile(newSqlLogFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				fmt.Println("Failed to open new log file")
			}

			// Update the logger to use the new log file
			//logger = zerolog.New(newLogFile).With().Timestamp().Logger()

			// Assign the new log file to the logFile variable
			logFile = newLogFile
			sqlLogFile = newSqlLogFile
		}
	}()

	log.Fatal(http.ListenAndServe(":8000", router))
}
func main() {
	RunMetrics()

}
