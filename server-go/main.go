package main

import (
	"fmt"
	"net/http"
)

func fortuneRequestHandler(w http.ResponseWriter, r *http.Request) {

	// Only handle GET requests
	// and return a 405 Method Not Allowed for other methods
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func main() {
	fmt.Println("Hello, World!")
}
