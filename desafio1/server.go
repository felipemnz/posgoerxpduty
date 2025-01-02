package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type CurrencyResponse struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

func fetchDollarRate(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var currency CurrencyResponse
	err = json.NewDecoder(resp.Body).Decode(&currency)
	if err != nil {
		return "", err
	}
	return currency.USDBRL.Bid, nil
}

func saveRateToDB(ctx context.Context, db *sql.DB, rate string) error {
	query := "INSERT INTO rates (rate, timestamp) VALUES (?, ?)"
	_, err := db.ExecContext(ctx, query, rate, time.Now())
	return err
}

func handleCotacao(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancel()

	rate, err := fetchDollarRate(ctx)
	if err != nil {
		http.Error(w, "Failed to fetch dollar rate", http.StatusInternalServerError)
		log.Println("Error fetching rate:", err)
		return
	}

	dbCtx, dbCancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer dbCancel()

	db, err := sql.Open("sqlite3", "./rates.db")
	if err != nil {
		log.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	err = saveRateToDB(dbCtx, db, rate)
	if err != nil {
		log.Println("Error saving rate to DB:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"bid": rate})
}

func main() {
	db, err := sql.Open("sqlite3", "./rates.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS rates (id INTEGER PRIMARY KEY, rate TEXT, timestamp DATETIME)")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/cotacao", handleCotacao)
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}