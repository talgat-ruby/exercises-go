package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/data"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type application struct {
	config config
	logger *log.Logger
	models data.Models
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://postgres:postgres@localhost/bloggingdb?sslmode=disable", "PostgreSQL DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	logger.Printf("database connection is established")

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	fmt.Println("app", app.routes())

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
	//ctx, cancel := context.WithCancel(context.Background())

	// // db
	// _, err := db.New()
	// if err != nil {
	// 	slog.ErrorContext(
	// 		ctx,
	// 		"initialize service error",
	// 		"service", "db",
	// 		"error", err,
	// 	)
	// 	panic(err)
	// }

	// // api
	// a := api.New()
	// if err := a.Start(ctx); err != nil {
	// 	slog.ErrorContext(
	// 		ctx,
	// 		"initialize service error",
	// 		"service", "api",
	// 		"error", err,
	// 	)
	// 	panic(err)
	// }

	// go func() {
	// 	shutdown := make(chan os.Signal, 1)   // Create channel to signify s signal being sent
	// 	signal.Notify(shutdown, os.Interrupt) // When an interrupt is sent, notify the channel

	// 	sig := <-shutdown
	// 	slog.WarnContext(ctx, "signal received - shutting down...", "signal", sig)

	// 	cancel()
	// }()
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
