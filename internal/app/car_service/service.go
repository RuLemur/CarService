package car_service

import (
	"context"
	"github.com/RuLemur/CarService/internal/app/datastruct"
	"github.com/RuLemur/CarService/internal/queue"
	"github.com/RuLemur/CarService/internal/repo"
)

type Service struct {
	db          *repo.QueryLogger
	queueClient *queue.Client
}

func NewService(db *repo.QueryLogger, queueClient *queue.Client) *Service {
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

func (s *Service) AddCar(ctx context.Context, userId int64, car *datastruct.UserCar) (int64, error) {
	err := repo.AddCar(s.db, userId, car)
	if err != nil {
		return 0, err
	}
	return car.ID, nil
}

func (s *Service) GetUserCars(ctx context.Context, userId int64) ([]*datastruct.UserCar, error) {
	cars, err := repo.GetUserCars(s.db, userId)
	if err != nil {
		return nil, err
	}
	return cars, nil
}