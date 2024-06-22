package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Request struct to map the JSON request
type Request struct {
	ID   string  `json:"id"`
	Cost float64 `json:"COST"`
}

// Response struct to map the JSON response
type Response struct {
	ID   string `json:"id"`
	Resp string `json:"RESP"`
}

// handler function to handle the POST request
func handler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request

	// Decode the incoming JSON request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Create the response
	resp := Response{
		ID:   req.ID,
		Resp: "SUCCESS",
	}

	// Encode the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/api", handler)

	log.Println("Server starting on port 443 with HTTPS...")
	// Provide the paths to your certificate and key files
	log.Fatal(http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil))
}
