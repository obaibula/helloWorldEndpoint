package main

import (
	"errors"
	"time"
)

type date time.Time

func (d *date) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return errors.New("date.UnmarshalJSON: input is not a JSON string")
	}
	data = data[len(`"`) : len(data)-len(`"`)]
	var err error
	parsedDate, err := time.Parse(time.DateOnly, string(data))
	*d = date(parsedDate)
	return err
}

func (d *date) toTime() time.Time {
	return time.Time(*d)
}
