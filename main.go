package main

import (
	"log"
	"net/http"

	"github.com/pteus/books-api/internal/handlers"
)

func main() {
	router := http.NewServeMux()

	handlers.SetupAuthRoutes(router)

	log.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
