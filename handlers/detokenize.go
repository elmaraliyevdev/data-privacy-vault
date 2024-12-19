package handlers

import (
	"data-privacy-vault/models"
	"data-privacy-vault/storage"
	"data-privacy-vault/utils"
	"encoding/json"
	"net/http"
)

func DetokenizeHandler(w http.ResponseWriter, r *http.Request) {
	var token models.Token
	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	storage.Store.Lock()
	defer storage.Store.Unlock()

	detokenizedData := make(map[string]map[string]interface{})
	flatData := make(map[string]string)

	for key, tokenValue := range token.Data {
		// Check if the token exists in storage
		if storedValue, exists := storage.Store.Data[tokenValue]; exists {
			// Decrypt the stored value
			decryptedValue, err := utils.Decrypt(storedValue["value"])
			if err != nil {
				http.Error(w, "Decryption failed", http.StatusInternalServerError)
				return
			}

			// Add to both detokenizedData (verbose) and flatData (for response)
			detokenizedData[key] = map[string]interface{}{
				"found": true,
				"value": decryptedValue,
			}
			flatData[key] = decryptedValue
		} else {
			detokenizedData[key] = map[string]interface{}{
				"found": false,
				"value": "",
			}
			flatData[key] = "" // Default empty string for unmatched tokens
		}
	}

	response := models.Token{
		ID:   token.ID,
		Data: flatData, // Only the flattened map[string]string is returned
	}
	json.NewEncoder(w).Encode(response)
}