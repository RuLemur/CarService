package endpoint

import (
	"car_service/internal/app/car_service"
	"car_service/internal/app/datastruct"
	"context"
	"fmt"
)

type GRPCRouter struct {
	s *car_service.Service
}

func NewGRPCRouter(s *car_service.Service) *GRPCRouter {
	return &GRPCRouter{s}
}

func (g *GRPCRouter) AddUser(ctx context.Context, request *AddUserRequest) (*AddUserResponse, error) {
	userID, err := g.s.AddUser(ctx, datastruct.User{Username: request.Username})
	if err != nil {
		return nil, err
	}
	return &AddUserResponse{
		Id: userID,
	}, nil
}

func (g *GRPCRouter) GetUser(ctx context.Context, request *GetUserRequest) (*GetUserResponse, error) {
	user, err := g.s.GetUser(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &GetUserResponse{
		Id:        user.ID,
		Username:  user.Username,
		GarageId:  user.GarageID.Int64,
		UpdatedAt: user.RegistrationAt,
	}, nil
}

func (g *GRPCRouter) CreateGarage(ctx context.Context, request *EmptyRequest) (*CreateGarageResponse, error) {
	garageId, err := g.s.CreateGarage(ctx)
	if err != nil {
		return nil, err
	}
	return &CreateGarageResponse{
		Id: garageId.GarageID,
	}, nil
}

func (g *GRPCRouter) GetGarage(ctx context.Context, request *GetGarageRequest) (*GetGarageResponse, error) {
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
	return &GetGarageResponse{
		GarageId: garages[0].GarageID,
		Cars:     cars,
	}, nil
}

func (g *GRPCRouter) CarSearch(ctx context.Context, request *CarSearchRequest) (*CarSearchResponse, error) {
	carModel := datastruct.CarModel{
		Brand: request.Brand,
		Model: request.Model,
	}
	models, err := g.s.SearchCarModel(ctx, &carModel)
	if err != nil {
		return nil, err
	}
	var cars []*Car
	for _, model := range models {
		cars = append(cars, &Car{
			Id:         model.ID,
			Brand:      model.Brand,
			Model:      model.Model,
			Equipment:  model.Equipment,
			EngineType: model.EngineType,
		})
	}
	tests := &CarSearchResponse{
		Car: cars,
	}
	return tests, nil
}

func (g *GRPCRouter) AddToGarage(ctx context.Context, request *AddToGarageRequest) (*AddToGarageResponse, error) {
	car := datastruct.UserCar{
		GarageId: request.GarageId,
	}
	g.s.AddToGarage(ctx, &car)
	return &AddToGarageResponse{CarId: 0}, nil
}
