package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from a Go app in a scratch Docker image!")
}

func main() {
    fmt.Println("Server is starting on port 8080...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
