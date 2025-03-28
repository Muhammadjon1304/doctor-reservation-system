package handlers

import (
	"encoding/json"
	"net/http"

	"doctor-reservation-system/internal/auth"
)

type AuthHandler struct {
	service *auth.Service
}

func NewAuthHandler(service *auth.Service) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user auth.User

	// Parse request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Register user
	err = h.service.Register(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Parse request body
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Authenticate user
	user, err := h.service.Login(loginData.Username, loginData.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Respond with user info (in real app, you'd generate a token)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id":  user.ID,
		"username": user.Username,
	})
}
