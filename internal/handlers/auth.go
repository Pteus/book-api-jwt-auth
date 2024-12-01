package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/pteus/books-api/internal/repositories"
	"github.com/pteus/books-api/internal/services"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SetupAuthRoutes(router *http.ServeMux, db *sql.DB) {
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)

	router.HandleFunc("POST /register", func(w http.ResponseWriter, r *http.Request) {
		handleRegister(w, r, authService)
	})

	router.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
		handleLogin(w, r, authService)
	})
}

func handleRegister(w http.ResponseWriter, r *http.Request, authService services.AuthService) {
	req := new(AuthRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	err = authService.Register(req.Username, req.Password)
	if err != nil {
		http.Error(w, "error registering user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully - Please login"})
}

func handleLogin(w http.ResponseWriter, r *http.Request, authService services.AuthService) {
	req := new(AuthRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	token, err := authService.Login(req.Username, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
