package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	fmt.Println("Web app running on localhost:3000")

	http.HandleFunc("/readiness", func(w http.ResponseWriter, r *http.Request) {
		// Check if the application is ready to serve traffic (e.g., database connection is established)
		// Return a 200 OK if ready, otherwise return a non-2xx status code.
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/liveness", func(w http.ResponseWriter, r *http.Request) {
		// Check if the application is alive (e.g., by performing a simple health check)
		// Return a 200 OK if alive, otherwise return a non-2xx status code.
		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":3000", nil)

}
