package middleware

import (
	"net/http"
	"strings"
)

// AuthMiddleware ensures that requests are authenticated.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized: Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Check the token format and value
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" || !validateToken(parts[1]) {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		// If valid, pass control to the next handler
		next.ServeHTTP(w, r)
	})
}

// A dummy token validation function (replace with your own logic)
func validateToken(token string) bool {
	// In a real system, you might verify the token against a database or use JWT
	return token == "valid-token"
}