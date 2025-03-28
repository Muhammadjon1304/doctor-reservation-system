package handlers

import (
	"doctor-reservation-system/internal/doctors"
	"encoding/json"
	"net/http"
	"time"

	"doctor-reservation-system/internal/notifications"
	"doctor-reservation-system/internal/reservations"
	"doctor-reservation-system/pkg/database"
	"doctor-reservation-system/pkg/middleware"
)

type ReservationHandler struct {
	reservationService  *reservations.Service
	notificationService *notifications.Service
	doctorRepo          *doctors.Repository
}

func NewReservationHandler(
	reservationService *reservations.Service,
	notificationService *notifications.Service,
	db *database.DB, // Pass the database connection
) *ReservationHandler {
	return &ReservationHandler{
		reservationService:  reservationService,
		notificationService: notificationService,
		doctorRepo:          doctors.NewRepository(db.SQL), // Use the SQL connection
	}
}

func (h *ReservationHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

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

	// Get user from repository or context
	user, err := h.reservationService.GetUserByID(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Get doctor details using the repository
	doctor, err := h.doctorRepo.GetDoctorByID(reqBody.DoctorID)
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
