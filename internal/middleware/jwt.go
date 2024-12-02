package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/pteus/books-api/internal/utils"
)

func RequireJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			http.Error(w, "invalid token format", http.StatusUnauthorized)
			return
		}

		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid token: %v", err), http.StatusUnauthorized)
			return
		}

		username, ok := (*claims)["username"].(string)
		if !ok || username == "" {
			http.Error(w, "invalid token claims", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "username", username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
