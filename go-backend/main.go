// main.go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/data", handleData)

	fmt.Println("Server running on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

func handleData(w http.ResponseWriter, r *http.Request) {
	// Handle the /api/data route and return a response
	w.Write([]byte(`{"message": "Hello from Golang!"}`))
}
