package repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rulemur/car_service/internal/app/datastruct"
	"log"
)

type Repository interface {
}

func AddUser(db *QueryLogger, user *datastruct.User) error {
	err := db.QueryRowx(`INSERT INTO users (username) VALUES ($1) RETURNING id`, user.Username).Scan(&user.ID)
	return err
}

func GetUser(db *QueryLogger, userId int64) (*datastruct.User, error) {
	var user datastruct.User
	err := sqlx.Get(db, &user, `SELECT * FROM users where id = $1`, userId)
	return &user, err
}

func CreateGarage(db *QueryLogger, garage *datastruct.Garage) error {
	err := db.QueryRowx(`INSERT INTO garage DEFAULT VALUES RETURNING id`).Scan(&garage.GarageID)
	return err
}

func GetGarage(db *QueryLogger, garageId int64) ([]*datastruct.Garage, error) {
	var garages []*datastruct.Garage
	err := sqlx.Select(db, &garages, `SELECT * FROM garage where id = $1`, garageId)

	return garages, err
}

func SearchCarModel(db *QueryLogger, filter map[string]string, limit int64) ([]*datastruct.CarModel, error) {
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
	} else {
		err := sqlx.Select(db, &models, `SELECT * FROM car_models limit 10`)
		if err != nil {
			return nil, err
		}
		return models, nil
	}
	return []*datastruct.CarModel{}, nil
}

func AddToGarage(db *QueryLogger, userCar *datastruct.UserCar) {
	db.QueryRowx(`INSERT INTO garage (id, car_id) VALUES ($1,$2) RETURNING id`, userCar.GarageId, userCar.CarId)
}
