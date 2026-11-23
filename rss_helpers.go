package main

import (
	"database/sql"
	"fmt"
	"time"
)

func validateTime(pubDate string) (sql.NullTime, error) {
	var parsedTime sql.NullTime
	parsed, err := time.Parse(time.RFC1123, pubDate)
	if err == nil {
		parsedTime.Time = parsed
		parsedTime.Valid = true
		return parsedTime, nil
	} else {
		fmt.Println(err)
	}
	parsed, err = time.Parse(time.RFC1123Z, pubDate)
	if err == nil {
		parsedTime.Time = parsed
		parsedTime.Valid = true
		return parsedTime, nil
	} else {
		fmt.Println(err)
	}
	parsedTime.Time = time.Now()
	parsedTime.Valid = false

	return parsedTime, err
}

/*
func validateDesc(desc string) (sql.NullString, error) {
	var parsedDesc sql.NullString
	if desc == ""
}
i*/
