package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func LogingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Println("Start at ", start)
		next.ServeHTTP(w, r)
		fmt.Println("End at ", time.Since(start))
	})
}
