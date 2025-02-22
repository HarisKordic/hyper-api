package utils

import (
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// CorsMiddleware sets the CORS headers
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

var (
	visitors = make(map[string]*rate.Limiter)
	mu       sync.RWMutex
)

func RateLimitMiddleware(next http.Handler) http.Handler {
	// 100 requests per second, burst of 500
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Use IP or token as identifier
		ip := r.RemoteAddr

		mu.Lock()
		if _, exists := visitors[ip]; !exists {
			visitors[ip] = rate.NewLimiter(rate.Every(time.Second/100), 500)
		}
		mu.Unlock()

		if !visitors[ip].Allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
