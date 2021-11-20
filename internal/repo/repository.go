package repo

import (
	"fmt"
	"github.com/RuLemur/CarService/internal/app/datastruct"
	"github.com/jmoiron/sqlx"
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


func AddCar(db *QueryLogger, userId int64, car *datastruct.UserCar) error {
	err := db.QueryRowx(`INSERT INTO user_car (user_id, model_id, production_year, mileage, car_name) VALUES ($1,$2,$3,$4,$5) RETURNING id`,
		userId,
		car.ModelID,
		car.Year,
		car.Mileage,
		car.CarName).Scan(&car.ID)
	return err
}


func GetUserCars(db *QueryLogger, userId int64) ([]*datastruct.UserCar, error) {
	var cars []*datastruct.UserCar
	err := sqlx.Select(db, &cars, `SELECT * FROM user_car where user_id = $1`, userId)
	return cars, err
}