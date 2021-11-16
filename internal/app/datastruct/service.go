package datastruct

import "time"

type GarageInfo struct {
	GarageName string
}

type User struct {
	ID             int64     `db:"id"`
	Username       string    `db:"username"`
	CarID          int64     `db:"car_id"`
	RegistrationAt time.Time `db:"registration_at"`
}
