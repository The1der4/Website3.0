package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Received request to /api/hello")
    response := Response{Message: "Hello from Go backend!"}
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
    log.Println("Sent response from /api/hello")
}

func main() {
	// Define API routes
	mux := http.NewServeMux()
	mux.HandleFunc("/api/hello", apiHandler)

	// Wrap with CORS middleware
	handler := enableCORS(mux)

	// Start server
	log.Println("Starting Go backend server on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
