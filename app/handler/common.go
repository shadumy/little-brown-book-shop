package handler

import (
	"encoding/json"
	"go-with-compose/config"
	"net/http"
)

var jwtKey []byte = config.GetConfig().Auth.Jwtkey
var systemUser string = config.GetConfig().Auth.Username
var systemPassword string = config.GetConfig().Auth.Password

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}
