package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type Booking struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	PlaceID int    `json:"place_id"`
	From    string `json:"from"`
	To      string `json:"to"`
}

type BooklistResponse struct {
	Bookings []Booking `json:"bookings"`
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func main() {
	port := flag.String("port", "8080", "port to listen on")
	flag.Parse()

	host := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	pass := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "contest")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, dbPort, user, pass, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open DB connection: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	mux.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		q := r.URL.Query()
		placeID := q.Get("place_id")
		userID := q.Get("user_id")
		fromStr := q.Get("from")
		toStr := q.Get("to")

		if placeID == "" || userID == "" || fromStr == "" || toStr == "" {
			http.Error(w, "Missing required query parameters", http.StatusBadRequest)
			return
		}

		timeFrom, err := time.Parse(time.RFC3339, fromStr)
		if err != nil {
			http.Error(w, "Invalid 'from' format", http.StatusBadRequest)
			return
		}

		timeTo, err := time.Parse(time.RFC3339, toStr)
		if err != nil {
			http.Error(w, "Invalid 'to' format", http.StatusBadRequest)
			return
		}

		query := `
			INSERT INTO bookings (place_id, user_id, time_from, time_to)
			SELECT $1, $2, $3, $4
			WHERE NOT EXISTS (
				SELECT 1 FROM bookings
				WHERE place_id = $1 AND time_from < $4 AND time_to > $3
			) RETURNING id;
		`

		var newID int
		err = db.QueryRow(query, placeID, userID, timeFrom.UTC(), timeTo.UTC()).Scan(&newID)

		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusConflict)
				return
			}
			log.Printf("DB insert error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/booklist", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		q := r.URL.Query()
		userID := q.Get("user_id")
		placeID := q.Get("place_id")

		var rows *sql.Rows
		var err error

		if userID != "" {
			rows, err = db.Query("SELECT id, user_id, place_id, time_from, time_to FROM bookings WHERE user_id = $1 ORDER BY time_from ASC, id ASC", userID)
		} else if placeID != "" {
			rows, err = db.Query("SELECT id, user_id, place_id, time_from, time_to FROM bookings WHERE place_id = $1 ORDER BY time_from ASC, id ASC", placeID)
		} else {
			http.Error(w, "Must provide user_id or place_id", http.StatusBadRequest)
			return
		}

		if err != nil {
			log.Printf("DB query error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		bookings := make([]Booking, 0)

		for rows.Next() {
			var b Booking
			var tFrom, tTo time.Time

			if err := rows.Scan(&b.ID, &b.UserID, &b.PlaceID, &tFrom, &tTo); err != nil {
				log.Printf("Row scan error: %v", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			b.From = tFrom.UTC().Format(time.RFC3339)
			b.To = tTo.UTC().Format(time.RFC3339)
			bookings = append(bookings, b)
		}

		resp := BooklistResponse{Bookings: bookings}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	addr := fmt.Sprintf(":%s", *port)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
