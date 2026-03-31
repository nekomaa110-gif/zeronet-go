package handlers

import (
	"encoding/json"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := map[string]string{
		"status":  "ok",
		"message": "ZERO NET BACKEND",
	}
	json.NewEncoder(w).Encode(data)
}