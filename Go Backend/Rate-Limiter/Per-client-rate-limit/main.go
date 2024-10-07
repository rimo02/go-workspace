package main

import (
	"encoding/json"
	"golang.org/x/time/rate"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func perClientRateLimiter(next func(w http.ResponseWriter, r *http.Request)) http.Handler {

	type Client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}
	var (
		mu      sync.Mutex
		clients = make(map[string]*Client)
	)
	go func() {
		for {
			time.Sleep(time.Minute)
			mu.Lock()
			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}
			mu.Unlock()
		}
	}()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		mu.Lock()
		if _, found := clients[ip]; !found {
			clients[ip] = &Client{
				limiter: rate.NewLimiter(2, 6),
			}
		}
		clients[ip].lastSeen = time.Now()
		if !clients[ip].limiter.Allow() {
			mu.Unlock()
			message := Message{
				Status: "Request Failed",
				Body:   "The API is at capacity, try again later.",
			}

			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(&message)
			return
		}
		mu.Unlock()
		next(w, r)
	})
}

func endpointHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := Message{
		Status: "Successfull",
		Body:   "Hi, You have rwached the API",
	}
	err := json.NewEncoder(w).Encode(&message)
	if err != nil {
		return
	}
}
func main() {
	http.Handle("/ping", perClientRateLimiter((endpointHandler)))
	err := http.ListenAndServe(":9080", nil)
	if err != nil {
		log.Println("There was an error listening on port 8080: ", err)
	}
}
