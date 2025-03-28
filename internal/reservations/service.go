package reservations

import (
	"errors"
	"time"

	"doctor-reservation-system/internal/auth"
	"doctor-reservation-system/internal/doctors"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateReservation(
	user *auth.User,
	doctor *doctors.Doctor,
	reservationTime time.Time,
) (*Reservation, error) {
	// Validate input
	if user == nil || doctor == nil {
		return nil, errors.New("invalid user or doctor")
	}

	// Check if reservation time is in the future
	if reservationTime.Before(time.Now()) {
		return nil, errors.New("reservation time must be in the future")
	}

	// Create reservation
	reservation := &Reservation{
		User:            user,
		Doctor:          doctor,
		ReservationTime: reservationTime,
		Status:          "scheduled",
	}

	err := s.repo.CreateReservation(reservation)
	if err != nil {
		return nil, err
	}

	return reservation, nil
}

func (s *Service) GetUserReservations(userID int) ([]Reservation, error) {
	return s.repo.GetUserReservations(userID)
}
