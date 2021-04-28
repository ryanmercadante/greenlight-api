package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Retrieve the "id" URL parameter from the current request context, then convert
// it to an integer and return it. If the operation isn't successful, return 0
// and an error.
func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

// Define a writeJSON() helper for sending responses. This takes the destination
// http.ResponseWriter, the HTTP status code to send, the data to encode to JSON,
// and a header map containing any additional HTTP headers we want to include in
// the response.
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	// Encode the data to JSON, returning the error if there was one.
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Append a newline to make it easier to view in terminal applications
	js = append(js, '\n')

	// Add any headers that we want to include. We loop through the header map and
	// add each header to the http.ResponseWriter header map. Note that it's OK if
	// the provided header map is nil. Go doesn't throw an error if you try to range
	// over (or generally, read from) a nil map.
	for key, value := range headers {
		w.Header()[key] = value
	}

	// Add the "Content-Type: application/json" header, then write the status code
	// and JSON response.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
