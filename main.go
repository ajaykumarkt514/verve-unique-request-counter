package main

import (
	"log"
	"net/http"
	"os"
	"verve-unique-request-counter/internal/middleware"

	"verve-unique-request-counter/internal/metric"

	uniquerequest "verve-unique-request-counter/internal/handler/unique-request"
)

func main() {
	// Open (or create if missing) "requests.log" with append and write permissions
	logFile, err := os.OpenFile("requests.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		// Log and exit if there’s an error opening the log file
		log.Fatalf("Failed to open log file: %v", err)
	}

	// Close the log file when the main function exits
	defer logFile.Close()

	// Direct all log output to the specified log file instead of the console
	log.SetOutput(logFile)

	// Initialize a new unique request handler
	handler := uniquerequest.New()

	// Wrap the main handler with the middleware
	http.Handle("/api/verve/accept", middleware.WithUniqueRequestLogging(http.HandlerFunc(handler.Get)))

	// Start a concurrent goroutine to log unique request counts periodically
	go metric.LogUniqueRequestCount()

	// Start HTTP server on port 8000 and handle incoming requests
	http.ListenAndServe(":8000", nil)
}
