package middleware

import (
	"net/http"
	"sync"
	"verve-unique-request-counter/internal/metric"
)

// Initialize middleware with a sync.Once to ensure LogUniqueRequestCount runs only once
var once sync.Once

// Middleware function that starts the goroutine and wraps the handler
func WithUniqueRequestLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Start the logging goroutine once
		once.Do(func() {
			go metric.LogUniqueRequestCount()
		})

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
