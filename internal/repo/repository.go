package repo

import (
	"car-service/internal/app/datastruct"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type Repository interface {
}

func AddUser(db *sqlx.DB, user *datastruct.User) error {
	err := db.QueryRowx(`INSERT INTO users (username) VALUES ($1) RETURNING id`, user.Username).Scan(&user.ID)
	return err
}

func GetUser(db *sqlx.DB, userId int64) (*datastruct.User, error) {
	var user datastruct.User
	err := db.Get(&user, `SELECT * FROM users where id = $1`, userId)
	return &user, err
}

func CreateGarage(db *sqlx.DB, garage *datastruct.Garage) error {
	err := db.QueryRowx(`INSERT INTO garage DEFAULT VALUES RETURNING id`).Scan(&garage.GarageID)
	return err
}

func GetGarage(db *sqlx.DB, garageId int64) ([]*datastruct.Garage, error) {
	var garages []*datastruct.Garage
	err := db.Select(&garages, `SELECT * FROM garage where id = $1`, garageId)

	return garages, err
}

func SearchCarModel(db *sqlx.DB, filter map[string]string, limit int64) ([]*datastruct.CarModel, error) {
	var models []*datastruct.CarModel

	var filterString string
	for k, v := range filter {
		if v != "" {
			if filterString != "" {
				filterString = fmt.Sprintf("%s AND %s = '%s'", filterString, k, v)
			} else {
				filterString = fmt.Sprintf("%s = '%s'", k, v)
			}
		}
	}
	fmt.Println(limit)
	if filterString != "" {
		rows, err := db.Queryx(fmt.Sprintf("SELECT * FROM car_models WHERE %s", filterString))
		for rows.Next() {
			var model datastruct.CarModel
			err := rows.StructScan(&model)
			if err != nil {
				log.Fatalln(err)
			}
			models = append(models, &model)
		}
		if err != nil {
			return nil, err
		}

		return models, nil
	}
	return []*datastruct.CarModel{}, nil
}

func AddToGarage(db *sqlx.DB, userCar *datastruct.UserCar) {
	db.MustExec(`INSERT INTO garage (id, car_id) VALUES ($1,$2) RETURNING id`, userCar.GarageId, userCar.CarId)
}
