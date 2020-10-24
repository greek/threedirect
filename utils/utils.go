package utils

import (
	"net/http"
)

// GetIP gets the user IP from X-Forwarded-For header
// and returns the address.
func GetIP(r *http.Request) string {
	df := r.Header.Get("X-Forwarded-For")
	if df != "" {
		return df
	}
	return r.RemoteAddr
}

// EnableCors allows all origins to access the endpoint.
func EnableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

// SetJSON sets content type to json
func SetJSON(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
}
