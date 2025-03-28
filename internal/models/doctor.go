package models

import "time"

type Doctor struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	Name             string    `json:"name" gorm:"not null"`
	Specialty        string    `json:"specialty" gorm:"not null"`
	WorkingHourStart time.Time `json:"working_hour_start"`
	WorkingHourEnd   time.Time `json:"working_hour_end"`
}
