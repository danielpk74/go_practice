package utils

import (
	"books-list/models"
	"encoding/json"
	"net/http"
)

func sendErrors(w http.ResponseWriter, status int, err models.Error) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(err)
}

func sendSuccesss(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
