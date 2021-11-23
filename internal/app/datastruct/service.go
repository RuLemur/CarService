package datastruct

import (
	"database/sql"
	"time"
)

type Garage struct {
	GarageID int64         `db:"id"`
	CarID    sql.NullInt64 `db:"car_id"`
}

type User struct {
	ID             int64         `db:"id"`
	Username       string        `db:"username"`
	GarageID       sql.NullInt64 `db:"garage_id"`
	RegistrationAt time.Time     `db:"registration_at"`
}

type UserCar struct {
	ID      int64     `db:"id"`
	UserID  int64     `db:"user_id"`
	ModelID int64     `db:"model_id"`
	Year    int64     `db:"production_year"`
	Mileage int64     `db:"mileage"`
	CarName string    `db:"car_name"`
	AddedAt time.Time `db:"added_at"`
}

type CarModel struct {
	ID         string         `db:"id"`
	Brand      string         `db:"brand"`
	Model      string         `db:"model"`
	Equipment  sql.NullString `db:"equipment"`
	EngineType sql.NullString `db:"engine_type"`
	YearFrom   int64          `db:"year_from"`
	YearTo     sql.NullInt64  `db:"year_to"`
	ImgLink    sql.NullString `db:"img_link"`
	ModelLink  sql.NullString `db:"model_link"`
}
