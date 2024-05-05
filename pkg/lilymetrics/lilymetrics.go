package lilymetrics

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

type Event struct {
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	Details   string `json:"details"`
}


var db *sql.DB

// handleEvent listens for incoming event data and logs it into the PostgreSQL database.

func handleJobEvent(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var event Event
	// var event map[string]interface{}
	err := decoder.Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
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

	event.Timestamp = time.Now().UTC().Format(time.RFC3339)
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

	// Now you can use the event map
	fmt.Println("details:", event.Details)
	// defer r.Body.Close()

	// r = fmt.Sprintf(`{"Type":"%s","Details":"%s"}`, event.Type, event.Details)
	fmt.Println(r)

	sql := fmt.Sprintf(`"INSERT INTO logs (type, timestamp, details) VALUES ("%s", "%s","%s")"`, event.Type, event.Timestamp, event.Details)
	fmt.Println(sql)
	_, err = db.Exec("INSERT INTO logs (type, timestamp, details) VALUES ($1, $2, $3)", event.Type, event.Timestamp, event.Details)

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
func RunMetrics() {

	dbHost := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := "postgres" //os.Getenv("POSTGRES_DB")
	connStr := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"
	fmt.Println("connStr: ", connStr)
	MigrateUp()
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

	log.Fatal(http.ListenAndServe(":8000", router))
}
func main() {
	RunMetrics()
}
