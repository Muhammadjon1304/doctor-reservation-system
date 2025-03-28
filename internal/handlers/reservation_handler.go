package handlers

import (
	"doctor-reservation-system/internal/doctors"
	"encoding/json"
	"net/http"
	"time"

	"doctor-reservation-system/internal/auth"
	"doctor-reservation-system/internal/notifications"
	"doctor-reservation-system/internal/reservations"
)

type ReservationHandler struct {
	reservationService  *reservations.Service
	notificationService *notifications.Service
}

func NewReservationHandler(
	reservationService *reservations.Service,
	notificationService *notifications.Service,
) *ReservationHandler {
	return &ReservationHandler{
		reservationService:  reservationService,
		notificationService: notificationService,
	}
}

func (h *ReservationHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Get user from context (in real app, you'd use middleware)
	user := &auth.User{ID: 1} // Placeholder

	// Parse request body
	var reqBody struct {
		DoctorID        int       `json:"doctor_id"`
		ReservationTime time.Time `json:"reservation_time"`
	}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Get doctor details
	doctor, err := doctors.NewRepository(nil).GetDoctorByID(reqBody.DoctorID)
	if err != nil {
		http.Error(w, "Doctor not found", http.StatusNotFound)
		return
	}

	// Create reservation
	reservation, err := h.reservationService.CreateReservation(
		user,
		doctor,
		reqBody.ReservationTime,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Send notification
	go h.notificationService.SendReservationNotification(
		user.Phone,
		reservation,
	)

	// Respond with reservation details
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reservation)
}

func (h *ReservationHandler) List(w http.ResponseWriter, r *http.Request) {
	// Get user from context (in real app, you'd use middleware)
	userID := 1 // Placeholder

	// Get user's reservations
	reservations, err := h.reservationService.GetUserReservations(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with reservations
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reservations)
}
