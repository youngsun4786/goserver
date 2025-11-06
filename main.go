package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"

	// Create the multiplexer
	mux := http.NewServeMux();

	// Create the server with configuration
	server := &http.Server {
		Addr: ":" + port,
		Handler: mux,
	}

	// Start the server
	log.Printf("Serving on port %s\n", port)
	// ListenAndServe returns error
	log.Fatal(server.ListenAndServe())
}
