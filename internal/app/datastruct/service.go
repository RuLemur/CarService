package datastruct

import "time"

type Garage struct {
	GarageID int64
	CarIDs   []int64
}

type User struct {
	ID             int64     `db:"id"`
	Username       string    `db:"username"`
	GarageID       int64     `db:"garage_id"`
	RegistrationAt time.Time `db:"registration_at"`
}

type UserCar struct {
	GarageId int64 `db:"id"`
	CarId    int64 `db:"car_id"`
}

type CarModel struct {
	ID         string `db:"id"`
	Brand      string `db:"brand"`
	Model      string `db:"model"`
	Equipment  string `db:"equipment"`
	EngineType string `db:"engine_type"`
}
