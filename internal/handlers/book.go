package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/pteus/books-api/internal/middleware"
	"github.com/pteus/books-api/internal/models"
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
	req := new(CreateBookRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if req.Title == "" || req.Author == "" || req.Genre == "" {
		http.Error(w, "Title, Author and Genre are required", http.StatusBadRequest)
		return
	}

	username, ok := r.Context().Value("username").(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	book := models.Book{
		Title:    req.Title,
		Author:   req.Author,
		Genre:    req.Genre,
		Username: username,
	}

	err := bookService.Create(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Book created"})
}

func handleGetAllBooksByUsername(w http.ResponseWriter, r *http.Request, bookService services.BookService) {
	username, ok := r.Context().Value("username").(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	books, err := bookService.GetAllByUsername(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(books)
}
