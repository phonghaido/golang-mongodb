package main

import (
	"log"
	"net/http"

	v1 "github.com/phonghaido/golang-mongodb/api/v1"
)

func main() {
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("Server listing on port 8080...")

	mux.HandleFunc("/v1/mongodb", v1.HandlePostMongoDB)
	mux.HandleFunc("/v1/in-memory", v1.HandleInMemory)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe: %v\n", err)
	}
}
