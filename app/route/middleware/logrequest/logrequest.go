package logrequest

import (
	"fmt"
	"net/http"
	"time"
)

// Handler will log the HTTP requests
func Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05 PM"), r.RemoteAddr, r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}
