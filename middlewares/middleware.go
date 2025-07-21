package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// start time
		start := time.Now()
		fmt.Printf("Request entered | PATH: %v | Method: %v\n", r.URL.Path, r.Method)

		// next function
		next.ServeHTTP(w, r)

		// end time
		fmt.Printf("Request completed | PATH: %v | Method: %v | Time: %v\n", r.URL.Path, r.Method, time.Since(start))
	})
}
