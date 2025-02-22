package main

import (
	"hyper-api/db"
	"hyper-api/server"
	"log"
	"net/http"
)

func main() {
	// Initialize the database
	db.Init()

	// Initialize the router
	router := server.NewRouter()

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
