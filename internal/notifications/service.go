package notifications

import (
	"log"

	"doctor-reservation-system/internal/reservations"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) SendReservationNotification(phone string, reservation *reservations.Reservation) {
	// In a real implementation, you'd integrate with an SMS/email service
	log.Printf(
		"Sending notification to %s for reservation with doctor %s at %v",
		phone,
		reservation.Doctor.Name,
		reservation.ReservationTime,
	)
}
