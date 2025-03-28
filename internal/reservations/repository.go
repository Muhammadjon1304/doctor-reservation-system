package reservations

import (
	"database/sql"
	"doctor-reservation-system/internal/auth"
	"doctor-reservation-system/internal/doctors"
	"errors"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateReservation(reservation *Reservation) error {
	// Check for conflicting reservations
	var count int
	err := r.db.QueryRow(`
		SELECT COUNT(*) FROM reservations 
		WHERE doctor_id = $1 AND reservation_time = $2
	`, reservation.Doctor.ID, reservation.ReservationTime).Scan(&count)

	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("time slot already booked")
	}

	query := `
		INSERT INTO reservations (user_id, doctor_id, reservation_time, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err = r.db.QueryRow(
		query,
		reservation.User.ID,
		reservation.Doctor.ID,
		reservation.ReservationTime,
		"scheduled",
	).Scan(&reservation.ID)

	return err
}

func (r *Repository) GetUserReservations(userID int) ([]Reservation, error) {
	query := `
		SELECT r.id, r.reservation_time, r.status,
			   d.id, d.name, d.specialty,
			   u.id, u.username, u.email
		FROM reservations r
		JOIN doctors d ON r.doctor_id = d.id
		JOIN users u ON r.user_id = u.id
		WHERE r.user_id = $1
		ORDER BY r.reservation_time
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []Reservation
	for rows.Next() {
		var res Reservation
		res.User = &auth.User{}
		res.Doctor = &doctors.Doctor{}

		err := rows.Scan(
			&res.ID,
			&res.ReservationTime,
			&res.Status,
			&res.Doctor.ID,
			&res.Doctor.Name,
			&res.Doctor.Specialty,
			&res.User.ID,
			&res.User.Username,
			&res.User.Email,
		)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, res)
	}

	return reservations, nil
}
