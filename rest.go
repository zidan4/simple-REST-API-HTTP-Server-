package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Define a response structure
type Response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Set response header to JSON
	w.Header().Set("Content-Type", "application/json")
	
	// Create a response object
	response := Response{Message: "Hello, World!"}
	
	// Encode the response as JSON and send it
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Handle HTTP request
	http.HandleFunc("/hello", helloHandler)

	// Start server on port 8080
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
