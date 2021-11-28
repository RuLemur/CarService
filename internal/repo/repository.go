package repo

import (
	"context"
	"fmt"
	"github.com/RuLemur/CarService/internal/app/datastruct"
	"github.com/RuLemur/CarService/internal/config"
	"github.com/RuLemur/CarService/internal/logger"
	"github.com/jmoiron/sqlx"
)

var log = logger.NewDefaultLogger()

type Repository struct {
	db *sqlx.DB
}

func (r *Repository) Init(ctx context.Context) error {
	cfg := config.GetInstance().GetConfig()
	log = logger.NewDefaultLogger()

	log.Infof("Connecting to database: %s", cfg.Database.DBHost)
	var err error
	r.db, err = sqlx.Connect("pgx", cfg.Database.DBHost)
	if err != nil {
		return err
	}

	log.Infof("Connected to database")
	return nil
}

func (r *Repository) Ping(ctx context.Context) error {
	log.Debugf("Ping database.")
	_, err := r.db.Query(`SELECT 1`)
	return err
}

func (r *Repository) Close() error {
	log.Infof("Close DB Connection.")
	return r.db.Close()
}

func (r *Repository) AddUser(user *datastruct.User) error {
	err := r.db.QueryRowx(`INSERT INTO users (username) VALUES ($1) RETURNING id`, user.Username).Scan(&user.ID)
	return err
}

func (r *Repository) GetUser(userId int64) (*datastruct.User, error) {
	var user datastruct.User
	err := r.db.Get(&user, `SELECT * FROM users where id = $1`, userId)
	return &user, err
}

func (r *Repository) SearchCarModel(filter map[string]string, limit int64) ([]*datastruct.CarModel, error) {
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
		rows, err := r.db.Queryx(fmt.Sprintf("SELECT * FROM car_models WHERE %s", filterString))
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
		err := r.db.Select(&models, `SELECT * FROM car_models limit 10`)
		if err != nil {
			return nil, err
		}
		return models, nil
	}
	return []*datastruct.CarModel{}, nil
}

func (r *Repository) AddCar(userId int64, car *datastruct.UserCar) error {
	err := r.db.QueryRowx(`INSERT INTO user_car (user_id, model_id, production_year, mileage, car_name) VALUES ($1,$2,$3,$4,$5) RETURNING id`,
		userId,
		car.ModelID,
		car.Year,
		car.Mileage,
		car.CarName).Scan(&car.ID)
	return err
}

func (r *Repository) GetUserCars(userId int64) ([]*datastruct.UserCar, error) {
	var cars []*datastruct.UserCar
	err := r.db.Select(&cars, `SELECT * FROM user_car where user_id = $1`, userId)
	return cars, err
}
