package main

import (
	"data-privacy-vault/handlers"
	"data-privacy-vault/middleware"
	"net/http"
)

func main() {
	// Use AuthMiddleware to secure the /tokenize and /detokenize endpoints
	http.Handle("/tokenize", middleware.AuthMiddleware(http.HandlerFunc(handlers.TokenizeHandler)))
	http.Handle("/detokenize", middleware.AuthMiddleware(http.HandlerFunc(handlers.DetokenizeHandler)))

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}