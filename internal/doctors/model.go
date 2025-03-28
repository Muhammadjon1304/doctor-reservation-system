package doctors

import "time"

type Doctor struct {
	ID               int
	Name             string
	Specialty        string
	WorkingHourStart time.Time
	WorkingHourEnd   time.Time
}
