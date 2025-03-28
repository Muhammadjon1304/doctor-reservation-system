package handlers

import (
	"encoding/json"
	"net/http"

	"doctor-reservation-system/internal/doctors"
)

type DoctorHandler struct {
	service *doctors.Service
}

func NewDoctorHandler(service *doctors.Service) *DoctorHandler {
	return &DoctorHandler{service: service}
}

func (h *DoctorHandler) Search(w http.ResponseWriter, r *http.Request) {
	// Get search query from request
	query := r.URL.Query().Get("q")

	// Search doctors
	doctors, err := h.service.SearchDoctors(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with doctors
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(doctors)
}
