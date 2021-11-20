package router

import (
	"context"
	"fmt"
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

func (g *GRPCRouter) CreateGarage(ctx context.Context, _ *endpoint.EmptyRequest) (*endpoint.CreateGarageResponse, error) {
	garageId, err := g.s.CreateGarage(ctx)
	if err != nil {
		return nil, err
	}
	return &endpoint.CreateGarageResponse{
		Id: garageId.GarageID,
	}, nil
}

func (g *GRPCRouter) GetGarage(ctx context.Context, request *endpoint.GetGarageRequest) (*endpoint.GetGarageResponse, error) {
	garages, err := g.s.GetGarage(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	if len(garages) == 0 {
		return nil, fmt.Errorf("not found garage")
	}

	var cars []int64
	for _, garage := range garages {
		cars = append(cars, garage.CarID.Int64)
	}
	return &endpoint.GetGarageResponse{
		GarageId: garages[0].GarageID,
		Cars:     cars,
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
	tests := &endpoint.CarSearchResponse{
		Car: cars,
	}
	return tests, nil
}

func (g *GRPCRouter) AddToGarage(ctx context.Context, request *endpoint.AddToGarageRequest) (*endpoint.AddToGarageResponse, error) {
	car := datastruct.UserCar{
		GarageId: request.GarageId,
	}
	g.s.AddToGarage(ctx, &car)
	return &endpoint.AddToGarageResponse{CarId: 0}, nil
}
