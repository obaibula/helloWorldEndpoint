package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(dst)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data any) error {
	marshal, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.WriteHeader(status)
	w.Write(marshal)
	return nil
}
