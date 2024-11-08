package metric

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

// UniqueRequests for in-memory tracking of unique IDs
var UniqueRequests sync.Map

// requestCounter represents the structure of the JSON payload sent to the endpoint
type requestCounter struct {
	UniqueCount int `json:"request_counter"`
}

// Mock function to count unique requests in sync.Map
func countUniqueRequests() int {
	count := 0

	UniqueRequests.Range(func(_, _ interface{}) bool {
		count++
		return true
	})

	return count
}

// LogUniqueRequestCount logs the unique request count every minute using a time.Ticker
func LogUniqueRequestCount() {
	// Create a ticker that ticks every 1 minute
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop() // Stop the ticker when the function exits

	for {
		select {
		// Wait for the ticker's next tick
		case <-ticker.C:
			// Count unique requests and reset the map
			uniqueCount := countUniqueRequests()
			UniqueRequests = sync.Map{} // Reset unique requests for the new minute

			// Log the count of unique requests in the past minute
			log.Printf("Unique request count in the past minute: %d\n", uniqueCount)
		}
	}
}

// SendCount sends the count of unique requests as a JSON payload to the specified endpoint via an HTTP POST request
func SendCount(endpoint string) {
	// Get the current count of unique requests
	uniqueCount := countUniqueRequests()

	// Create a JSON object from the unique count using the requestCounter struct
	data, _ := json.Marshal(requestCounter{UniqueCount: uniqueCount})

	// Send a POST request to the provided endpoint with JSON data
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(data))
	if err != nil {
		// Log an error if the POST request fails and exit the function
		log.Printf("Failed to send count to endpoint: %v", err)

		return
	}

	// Log the HTTP status code returned by the endpoint
	log.Printf("Response status from endpoint: %v", resp.StatusCode)
}
