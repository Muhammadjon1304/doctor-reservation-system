package doctors

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Search(query string) ([]Doctor, error) {
	sqlQuery := `
		SELECT id, name, specialty, working_hour_start, working_hour_end 
		FROM doctors 
		WHERE name ILIKE $1 OR specialty ILIKE $1
	`

	rows, err := r.db.Query(sqlQuery, fmt.Sprintf("%%%s%%", query))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doctors []Doctor
	for rows.Next() {
		var d Doctor
		err := rows.Scan(
			&d.ID,
			&d.Name,
			&d.Specialty,
			&d.WorkingHourStart,
			&d.WorkingHourEnd,
		)
		if err != nil {
			return nil, err
		}
		doctors = append(doctors, d)
	}

	return doctors, nil
}

func (r *Repository) GetDoctorByID(id int) (*Doctor, error) {
	doctor := &Doctor{}
	query := `
		SELECT id, name, specialty, working_hour_start, working_hour_end
		FROM doctors 
		WHERE id = $1
	`

	err := r.db.QueryRow(query, id).Scan(
		&doctor.ID,
		&doctor.Name,
		&doctor.Specialty,
		&doctor.WorkingHourStart,
		&doctor.WorkingHourEnd,
	)

	if err != nil {
		return nil, err
	}

	return doctor, nil
}
