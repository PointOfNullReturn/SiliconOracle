package main

import (
	"fmt"
	"log"
	"net/http"
)

// infoRequestHandler handles the /api/wisdom/info endpoint
func infoRequestHandler(w http.ResponseWriter, r *http.Request) {

	// Only handle GET requests
	// and return a 405 Method Not Allowed for other methods
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Return the info response as JSON
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"status": "ok", "message": "The SiliconeOracle is dispensing wisdom."}`)
}

// wisdomRequestHandler handles the /api/wisdom/dispense endpoint
func wisdomRequestHandler(w http.ResponseWriter, r *http.Request) {

	// Only handle GET requests
	// and return a 405 Method Not Allowed for other methods
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func main() {

	// Setup HTTP Handler for Info API Requst
	http.HandleFunc("/api/wisdom/info", infoRequestHandler)

	// Setup HTTP Handler for Wisdom dispense API Request
	//http.HandleFunc("/api/wisdom/dispense", wisdomRequestHandler)

	// Start the HTTP server
	port := "8080"
	fmt.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
