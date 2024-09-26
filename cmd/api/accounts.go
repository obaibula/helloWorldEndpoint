package main

import (
	"helloWorldEndpoint/internal/data"
	"net/http"
	"time"
)

func (app *application) createAccountHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Birthday  date   `json:"birthday"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		app.logger.Error(err.Error())
		return
	}

	account := &data.Account{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Birthday:  time.Time(input.Birthday),
	}

	err = app.models.Accounts.Insert(account)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		app.logger.Error(err.Error())
		return
	}

	err = app.writeJSON(w, http.StatusCreated, account)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		app.logger.Error(err.Error())
	}
}
