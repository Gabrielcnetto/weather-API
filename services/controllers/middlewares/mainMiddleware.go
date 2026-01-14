package middlewares

import (
	"encoding/json"
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

type Client struct {
	limiter *rate.Limiter
}

var clients = make(map[string]*Client)
var mu sync.Mutex

func getClientByIp(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()
	if client, exist := clients[ip]; exist {
		return client.limiter
	}
	limiter := rate.NewLimiter(5, 1)
	clients[ip] = &Client{limiter: limiter}
	return limiter
}

func MainMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userIp := r.RemoteAddr
		limiter := getClientByIp(userIp)
		if !limiter.Allow() {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			response := map[string]interface{}{
				"error": "Too many Requests",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		next.ServeHTTP(w, r)
	})
}
