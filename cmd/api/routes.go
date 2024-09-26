package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello/{$}", app.HelloWorld)
	mux.HandleFunc("GET /hello/{name}", app.HelloName)

	return mux
}
