package main

import (
	"fmt"
	"net/http"
	"time"
)

func (app *application) listenAndServe() error {
	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", app.cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.logger.Info("Starting a server on ", "addr", srv.Addr)
	return srv.ListenAndServe()
}
