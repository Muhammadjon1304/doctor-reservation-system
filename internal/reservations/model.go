package reservations

import (
	"time"

	"doctor-reservation-system/internal/auth"
	"doctor-reservation-system/internal/doctors"
)

type Reservation struct {
	ID              int
	User            *auth.User
	Doctor          *doctors.Doctor
	ReservationTime time.Time
	Status          string
}
