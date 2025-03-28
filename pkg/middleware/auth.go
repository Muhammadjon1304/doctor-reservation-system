package middleware

import (
	"context"
	"net/http"

	"doctor-reservation-system/internal/auth"
)

func AuthMiddleware(authService *auth.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get username/password from Authorization header
			username, password, ok := r.BasicAuth()
			if !ok {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Authenticate user
			user, err := authService.Login(username, password)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Add user to context
			ctx := context.WithValue(r.Context(), "user", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
