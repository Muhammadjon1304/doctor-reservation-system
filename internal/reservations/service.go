package reservations

import (
	"errors"
	"time"

	"doctor-reservation-system/internal/auth"
	"doctor-reservation-system/internal/doctors"
)

type Service struct {
	repo     *Repository
	authRepo *auth.Repository
}

func NewService(repo *Repository, authRepo *auth.Repository) *Service {
	return &Service{
		repo:     repo,
		authRepo: authRepo,
	}
}

func (s *Service) GetUserByID(userID int) (*auth.User, error) {
	// Implement a method to fetch user by ID
	query := `
		SELECT id, username, email, phone
		FROM users 
		WHERE id = $1
	`

	user := &auth.User{}
	err := s.authRepo.DB.QueryRow(query, userID).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Phone,
	)

	if err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
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
