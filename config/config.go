package config

import "net/http"

// Headers sets certain headers.
func Headers(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", "threedirector")
		next.ServeHTTP(w, r)
	})
}
