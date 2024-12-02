package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/pteus/books-api/internal/middleware"
	"github.com/pteus/books-api/internal/repositories"
	"github.com/pteus/books-api/internal/services"
)

type CreateBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

func SetupBookRoutes(router *http.ServeMux, db *sql.DB) {
	bookRepo := repositories.NewBookRepository(db)
	bookService := services.NewBookService(bookRepo)

	router.HandleFunc("POST /book", func(w http.ResponseWriter, r *http.Request) {
		middleware.RequireJWT(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handleCreateBook(w, r, bookService)
		})).ServeHTTP(w, r)
	})

	router.HandleFunc("GET /book", func(w http.ResponseWriter, r *http.Request) {
		middleware.RequireJWT(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handleGetAllBooksByUsername(w, r, bookService)
		})).ServeHTTP(w, r)
	})
}

func handleCreateBook(w http.ResponseWriter, r *http.Request, bookService services.BookService) {
	json.NewEncoder(w).Encode(map[string]string{"message": "createbook"})
}

func handleGetAllBooksByUsername(w http.ResponseWriter, r *http.Request, bookService services.BookService) {
}
