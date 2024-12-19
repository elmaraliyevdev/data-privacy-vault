package handlers

import (
	"data-privacy-vault/models"
	"data-privacy-vault/storage"
	"data-privacy-vault/utils"
	"encoding/json"
	"math/rand"
	"net/http"
)

func TokenizeHandler(w http.ResponseWriter, r *http.Request) {
	var token models.Token
	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	storage.Store.Lock()
	defer storage.Store.Unlock()

	tokenizedData := make(map[string]string)
	for key, value := range token.Data {
		// Encrypt the value
		encryptedValue, err := utils.Encrypt(value)
		if err != nil {
			http.Error(w, "Encryption failed", http.StatusInternalServerError)
			return
		}

		// Generate a token for the encrypted value
		tokenizedKey := generateToken()
		tokenizedData[key] = tokenizedKey

		// Store the encrypted value with the tokenized key
		storage.Store.Data[tokenizedKey] = map[string]string{"key": key, "value": encryptedValue}
	}

	token.Data = tokenizedData
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(token)
}

func generateToken() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}