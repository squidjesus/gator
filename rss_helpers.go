package main

import (
	"database/sql"
	"errors"
	"html"

	//"fmt"
	"strings"
	"time"
)

func validateTime(pubDate string) (sql.NullTime, error) {
	var t time.Time
	cleanedTime := strings.TrimSpace(pubDate)
	parsedTime := sql.NullTime{
		Time:  t,
		Valid: false,
	}
	layouts := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC3339,
		time.RFC3339Nano,
		time.RFC822,
		time.RFC822Z,
	}
	for _, layout := range layouts {
		if parsed, err := time.Parse(layout, cleanedTime); err == nil {
			parsedTime.Time = parsed
			parsedTime.Valid = true
			return parsedTime, nil
		}
	}
	return parsedTime, errors.New("Time not parsed, returning invalid")
}

func validateDesc(desc string) (sql.NullString, error) {
	parsedDesc := sql.NullString{
		String: "",
		Valid:  false,
	}
	if desc != "" {
		parsedDesc.String = html.UnescapeString(desc)
		parsedDesc.Valid = true
		return parsedDesc, nil
	}
	return parsedDesc, errors.New("No description provided")
}
