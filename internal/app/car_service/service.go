package car_service

import (
	"car-service/internal/app/datastruct"
	"car-service/internal/queue"
	"car-service/internal/repo"
	"context"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	db          *sqlx.DB
	queueClient *queue.Client
}

func NewService(db *sqlx.DB, queueClient *queue.Client) *Service {
	return &Service{db, queueClient}
}

func (s *Service) AddUser(ctx context.Context, user datastruct.User) (int64, error) {
	//err := s.queueClient.SendMessageToQueue("Hi!")
	err := repo.AddUser(s.db, &user)
	return user.ID, err
}

func (s *Service) GetUser(ctx context.Context, userId int64) (*datastruct.User, error) {
	user, err := repo.GetUser(s.db, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) CreateGarage(ctx context.Context) (datastruct.Garage, error) {
	var garage datastruct.Garage
	err := repo.CreateGarage(s.db, &garage)
	return garage, err
}

func (s *Service) GetGarage(ctx context.Context, garageID int64) ([]*datastruct.Garage, error) {
	garage, err := repo.GetGarage(s.db, garageID)
	if err != nil {
		return nil, err
	}

	return garage, nil
}

func (s *Service) SearchCarModel(ctx context.Context, carModel *datastruct.CarModel) ([]*datastruct.CarModel, error) {
	filter := map[string]string{
		"brand": carModel.Brand,
		"model": carModel.Model,
	}
	model, err := repo.SearchCarModel(s.db, filter, 10)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (s *Service) AddToGarage(ctx context.Context, userCar *datastruct.UserCar) {
	repo.AddToGarage(s.db, userCar)
}
