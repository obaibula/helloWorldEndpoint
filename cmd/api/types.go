package main

import "time"

type date time.Time

func (d *date) UnmarshalJSON(bytes []byte) error {
	dateStr := string(bytes[1 : len(bytes)-1])
	parsedDate, err := time.Parse(time.DateOnly, dateStr)
	if err != nil {
		return err
	}
	*d = date(parsedDate)
	return nil
}
