// Package helpers.
package utils

// Imports.
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// WriteResponse Helper function write a response to the client.
func WriteResponse(w http.ResponseWriter, status int, data interface{}) {
	// Set Content Type to Json.
	w.Header().Set("Content-Type", "application/json")

	// Write Status.
	w.WriteHeader(status)

	// Convert the data to json.
	jsonData, err := json.Marshal(data)

	// Check for Error.
	if err != nil {
		// Notify.
		fmt.Println("Error Marshalling Data: " + err.Error())

		// Exit.
		os.Exit(1)
	}

	// Write the data to the client.
	w.Write(jsonData)
}

// ReadRequestBody Helper function to read a request body as json.
func ReadRequestBody(r *http.Request, response interface{}) error {
	// Define request body and possible error.
	var body []byte
	var err error

	// Read the request body.
	body, err = ioutil.ReadAll(r.Body)

	// Check for Error.
	if err != nil {
		// Return Error.
		return err
	}

	// Unmarshal the body into the response.
	err = json.Unmarshal(body, response)

	// Check for Error.
	if err != nil {
		// Return Error.
		return err
	}

	// Return nil. (no Error)
	return nil
}

// CheckRequestMethod Helper function to check HTTP Request Method.
func CheckRequestMethod(r *http.Request, method string, w http.ResponseWriter) bool {
	// Check if the request method is the same as the method.
	if r.Method == method {
		// Return true.
		return true
	}

	// Check for Writer.
	if w != nil {
		// If writer is valid, we just write a response.
		WriteResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed")
	}

	// Return false.
	return false
}
