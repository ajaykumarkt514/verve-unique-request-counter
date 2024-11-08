package unique_request

import (
	"fmt"
	"net/http"
	"verve-unique-request-counter/internal/metric"
)

// Get is the HTTP handler method for the uniqueRequest struct, implementing the UniqueRequest interface.
// It processes incoming GET requests, validates query parameters, stores unique uniqueRequest IDs,
// and optionally makes an HTTP uniqueRequest to an external endpoint.
func (u uniqueRequest) Get(w http.ResponseWriter, r *http.Request) {
	// Parse and validate the "id" query parameter
	id := r.URL.Query().Get("id")
	if id == "" {
		// If "id" is missing, return a 400 Bad Request error with a descriptive message
		http.Error(w, "id query parameter is required", http.StatusBadRequest)
		return
	}

	// Store the unique uniqueRequest ID in the sync.Map for tracking unique requests
	metric.UniqueRequests.Store(id, true)

	// Check if the "endpoint" query parameter is provided
	endpoint := r.URL.Query().Get("endpoint")
	if endpoint != "" {
		// If "endpoint" is provided, make a concurrent HTTP request to the endpoint
		// to send the unique unique request count
		go metric.SendCount(endpoint)
	}

	// Respond with "ok" to indicate successful uniqueRequest processing
	fmt.Fprintln(w, "ok")
}
