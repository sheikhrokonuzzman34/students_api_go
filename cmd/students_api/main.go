package main

import (
	"net/http"
	"students-api/internal/config"
)

func main() {
	// Load configuration
	cfg := config.MustLoad()
	// database setup

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// start server
}
