package main

import (
	"database/sql"
	"flag"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"helloWorldEndpoint/internal/data"
	"log"
	"log/slog"
	"os"
	"time"
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
	models *data.Models
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	var cfg config
	flag.StringVar(&cfg.db.dsn, "dsn", os.Getenv("DB_DSN"), "Database DSN")
	flag.IntVar(&cfg.port, "port", 8080, "Port for the application")
	flag.Parse()

	logHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(logHandler)

	db, err := openDB(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()
	logger.Info("database connection pool established")

	app := application{
		cfg:    &cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	err = app.listenAndServe()
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(15 * time.Minute)

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
