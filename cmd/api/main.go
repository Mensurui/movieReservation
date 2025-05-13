package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	http_transport "github.com/Mensurui/movieReservation/internal/transport/http"
	_ "github.com/lib/pq"
)

type config struct {
	DatabaseDSN string
}

func main() {
	//Get the db cfg from os
	var cfg config
	cfg.DatabaseDSN = os.Getenv("movieReccServerDatabase")
	//Have a open postgresql here
	ctx := context.Background()
	db, err := openDB(ctx, cfg)
	if err != nil {
		log.Printf("Error opening the db: %v", err)
	}
	defer db.Close()
	//Pass it as a dependency to a gin router and get it
	router := http_transport.NewServer(db)
	srv := &http.Server{
		Addr:         "0.0.0.0:8090",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	//Start the server here
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("ListenAndServe failed: %v", err)
	}

}

func openDB(ctx context.Context, config config) (*sql.DB, error) {
	db, err := sql.Open("postgres", config.DatabaseDSN)
	if err != nil {
		return nil, fmt.Errorf("error opening the db: %v", err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
