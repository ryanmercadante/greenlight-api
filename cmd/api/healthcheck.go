package main

import (
	"encoding/json"
	"net/http"
)

// Declare a handler which writes a plain-text response which information about
// the application status, operating environment, and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Create a map which holds the information that we want to send in the response.
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	// Pass the map to the json.Marshal() function. This returns a []byte slice
	// containing the encoded JSON. If there was an error, we log it and send the client
	// a generic error message.
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}

	// Append a newline to the JSON. This is just a small nicety to make it easier
	// to view in terminal applications.
	js = append(js, '\n')

	// Set the "Content-Type: application/json" header on the response.
	w.Header().Set("Content-Type", "application/json")

	// Use w.Write() to send the []byte slice containing the JSON as the response body.
	w.Write(js)
}
