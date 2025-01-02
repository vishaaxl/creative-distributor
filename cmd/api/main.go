package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/vishaaxl/creative-distributer/internal/data"
)

type application struct {
	config config
	models data.Models
}

type config struct {
	port int
	env  string
	db   db
}

type db struct {
	dsn          string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  time.Duration
}

func main() {
	cfg := config{
		port: 4000,
		env:  "development",
		db: db{
			dsn:          "postgres://postgres:mysecretpassword@localhost/creative?sslmode=disable",
			maxOpenConns: 25,
			maxIdleConns: 25,
			maxIdleTime:  time.Minute * 15,
		},
	}

	db, err := openDB(cfg)
	if err != nil {
		log.Println(err, nil)
	}

	defer db.Close()
	log.Println("database connection established")

	app := &application{
		config: cfg,
		models: data.NewModels(db),
	}

	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/v1/users/sendOtp", app.sendOTPHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("starting %s server on port: %d", app.config.env, app.config.port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	db.SetConnMaxIdleTime(cfg.db.maxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
