package main

import (
	"flag"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
)

type config struct {
	port int
	db   struct {
		dsn string
	}
}

type application struct {
	cfg    *config
	logger *slog.Logger
}

func main() {
	var cfg config
	flag.StringVar(&cfg.db.dsn, "dsn", "", "Database DSN")
	flag.IntVar(&cfg.port, "port", 8080, "Port for the application")
	flag.Parse()

	logHandler := slog.NewJSONHandler(os.Stdout, nil)

	app := application{
		cfg:    &cfg,
		logger: slog.New(logHandler),
	}

	err := app.listenAndServe()
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}
}
