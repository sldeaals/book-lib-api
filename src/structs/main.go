package structs

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus" // Using logrus for more sophisticated logging
)

// Logger interface for logging
type Logger interface {
	Error(args ...interface{})
	Fatal(args ...interface{})
}

// DefaultLogger is the default logger using logrus
type DefaultLogger struct{}

func (l *DefaultLogger) Error(args ...interface{}) {
	logrus.Error(args...)
}

func (l *DefaultLogger) Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

// Log is a global variable for logging. This can be set to a custom logger during testing.
var Log Logger = &DefaultLogger{}

// HandleError handles errors by logging and sending an appropriate HTTP response.
func HandleError(w http.ResponseWriter, r *http.Request, err *Error) {
	// Log the error for debugging purposes.
	Log.Error(fmt.Sprintf("Error: %+v", err))

	// Send a JSON response with the error information.
	ResponseJSON(w, err.Code, Response{
		Code:   err.Code,
		Status: err.Message,
		Data:   nil,
	})
}

// ResponseJSON sends a JSON response with the given HTTP status code and data.
func ResponseJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// LoadEnv loads environment variables from a .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		Log.Fatal("Error loading .env file")
	}
}
