package main

import (
	"log"
	"net/http"

	"github.com/ghenoo/microservices-go/cmd/db"
	"github.com/ghenoo/microservices-go/cmd/handlers"
)


func main() {
	db.InitDB()
    handlers.SetupRoutes()

    log.Println("Server is starting on port 8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("Failed to start the server: %v", err)
    }
}