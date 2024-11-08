package unique_request

import "net/http"

// uniqueRequest struct is an empty struct that implements the UniqueRequest interface.
type uniqueRequest struct{}

// UniqueRequest interface defines a contract for handling requests.
type UniqueRequest interface {
	Get(w http.ResponseWriter, r *http.Request)
}

// New function is a constructor that returns an instance of UniqueRequest.
func New() UniqueRequest {
	return uniqueRequest{}
}
