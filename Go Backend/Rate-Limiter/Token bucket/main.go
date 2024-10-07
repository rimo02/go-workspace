package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

// writer http.ResponseWriter: This is used to write the response back to the client. It allows you to set headers, status codes, and the response body.

// request *http.Request: This represents the incoming HTTP request. It contains information such as the request method (GET, POST, etc.), URL, headers, and possibly a body.
func endpointHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := Message{
		Status: "Successfull",
		Body:   "Hi, You have rwached the API",
	}
	err := json.NewEncoder(w).Encode(&message); 
	if err != nil {
		return
	}
}
func main() {
	http.Handle("/ping", RateLimiter((endpointHandler)))
	err := http.ListenAndServe(":9080", nil)
	if err != nil {
		log.Println("There was an error listening on port 8080: ", err)
	}
}
