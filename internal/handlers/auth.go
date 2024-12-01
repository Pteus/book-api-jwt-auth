package handlers

import (
	"encoding/json"
	"net/http"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SetupAuthRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /register", func(w http.ResponseWriter, r *http.Request) {
		handleRegister(w, r)
	})

	router.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
		handleLogin(w, r)
	})
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "User logged in successfully"})
}
