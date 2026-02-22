package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Dimisz/store-api/go-api/internal/jsonlog"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *jsonlog.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 3001, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	app := &application{
		config: cfg,
		logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/healthcheck", app.healthcheckHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.PrintInfo(fmt.Sprintf("starting %s server on %s", cfg.env, srv.Addr), nil)
	err := srv.ListenAndServe()
	logger.PrintFatal(err, nil)
}
