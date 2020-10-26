package utils

import (
	"net/http"
	"os"
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

func CheckFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}
