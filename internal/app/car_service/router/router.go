package router

import (
	"context"
	"github.com/RuLemur/CarService/internal/app/car_service"
	"github.com/RuLemur/CarService/internal/app/datastruct"
	"github.com/RuLemur/CarService/pkg/endpoint"
)

type GRPCRouter struct {
	s *car_service.Service
}

func NewGRPCRouter(s *car_service.Service) *GRPCRouter {
	return &GRPCRouter{s}
}

func (g *GRPCRouter) AddCar(ctx context.Context, request *endpoint.AddCarRequest) (*endpoint.AddCarResponse, error) {
	userCar := &datastruct.UserCar{
		ModelID: request.ModelId,
		Year:    request.ProductionYear,
		Mileage: request.Mileage,
		CarName: request.CarName,
	}
	carID, err := g.s.AddCar(ctx, request.UserId, userCar)
	if err != nil {
		return nil, err
	}
	return &endpoint.AddCarResponse{
		UserCarId: carID,
	}, nil
}

func (g *GRPCRouter) GetUserCars(ctx context.Context, request *endpoint.GetUserCarsRequest) (*endpoint.GetUserCarsResponse, error) {
	userCars, err := g.s.GetUserCars(ctx, request.UserId)
	if err != nil {
		return nil, err
	}
	if len(userCars) == 0 {
		return &endpoint.GetUserCarsResponse{UserId: request.UserId}, nil
	}
	var cars []*endpoint.UserCar
	for _, userCar := range userCars {
		cars = append(cars, &endpoint.UserCar{
			Id:             userCar.ID,
			ModelId:        userCar.ModelID,
			CarName:        userCar.CarName,
			ProductionYear: userCar.Year,
			Mileage:        userCar.Mileage,
			AddedAt:        userCar.AddedAt,
		})
	}
	return &endpoint.GetUserCarsResponse{
		UserId:   request.UserId,
		UserCars: cars,
	}, nil
}

func (g *GRPCRouter) AddUser(ctx context.Context, request *endpoint.AddUserRequest) (*endpoint.AddUserResponse, error) {
	userID, err := g.s.AddUser(ctx, datastruct.User{Username: request.Username})
	if err != nil {
		return nil, err
	}
	return &endpoint.AddUserResponse{
		Id: userID,
	}, nil
}

func (g *GRPCRouter) GetUser(ctx context.Context, request *endpoint.GetUserRequest) (*endpoint.GetUserResponse, error) {
	user, err := g.s.GetUser(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &endpoint.GetUserResponse{
		Id:        user.ID,
		Username:  user.Username,
		GarageId:  user.GarageID.Int64,
		UpdatedAt: user.RegistrationAt,
	}, nil
}

func (g *GRPCRouter) CarSearch(ctx context.Context, request *endpoint.CarSearchRequest) (*endpoint.CarSearchResponse, error) {
	carModel := datastruct.CarModel{
		Brand: request.Brand,
		Model: request.Model,
	}
	models, err := g.s.SearchCarModel(ctx, &carModel)
	if err != nil {
		return nil, err
	}
	var cars []*endpoint.Car
	for _, model := range models {
		cars = append(cars, &endpoint.Car{
			Id:         model.ID,
			Brand:      model.Brand,
			Model:      model.Model,
			Equipment:  model.Equipment,
			EngineType: model.EngineType,
		})
	}
	response := &endpoint.CarSearchResponse{
		Car: cars,
	}
	return response, nil
}
