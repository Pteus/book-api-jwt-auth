package main

import (
	"log"
	"net/http"

	"github.com/pteus/books-api/internal/configs"
	"github.com/pteus/books-api/internal/handlers"
)

func main() {
	router := http.NewServeMux()

	db := configs.InitDatabase()
	defer db.Close()

	handlers.SetupAuthRoutes(router, db)
	handlers.SetupBookRoutes(router, db)

	log.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
