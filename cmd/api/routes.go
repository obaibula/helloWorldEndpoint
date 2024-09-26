package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello/{$}", app.helloWorld)
	mux.HandleFunc("GET /hello/{name}", app.helloName)
	mux.HandleFunc("POST /accounts", app.createAccountHandler)

	return mux
}
