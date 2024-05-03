package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Event struct {
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	Details   string `json:"details"`
}

var db *sql.DB

// handleEvent listens for incoming event data and logs it into the PostgreSQL database.
func handleEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	_, err := db.Exec("INSERT INTO events (type, timestamp, details) VALUES ($1, $2, $3)", event.Type, event.Timestamp, event.Details)
	if err != nil {
		log.Printf("Error inserting event into database: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("Event Received and Logged: %+v", event)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Event received and logged successfully"))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	connStr := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/metrics-dashboard/log", handleEvent).Methods("POST") // Only keep the event handler route

	log.Fatal(http.ListenAndServe(":8000", router))
}
