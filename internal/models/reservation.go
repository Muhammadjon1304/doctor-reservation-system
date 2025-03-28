package models

import "time"

type Reservation struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	UserID          uint      `json:"user_id"`
	User            User      `json:"user" gorm:"foreignKey:UserID"`
	DoctorID        uint      `json:"doctor_id"`
	Doctor          Doctor    `json:"doctor" gorm:"foreignKey:DoctorID"`
	ReservationTime time.Time `json:"reservation_time"`
	Status          string    `json:"status" gorm:"default:'scheduled'"`
}
