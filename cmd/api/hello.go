package main

import (
	"net/http"
)

func (app *application) helloWorld(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		app.logger.Error("failed to write in helloWorld handler", "cause", err.Error())
	}
}

func (app *application) helloName(w http.ResponseWriter, r *http.Request) {
	response := "hello, " + r.PathValue("name")
	_, err := w.Write([]byte(response))
	if err != nil {
		app.logger.Error("failed to write in helloWorld handler", "cause", err.Error())
	}
}
