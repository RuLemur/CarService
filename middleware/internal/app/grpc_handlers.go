package app

import (
	"car-service/internal/app/car_service/endpoint"
	"car-service/middleware/internal/datastruct"
)

func (a *Api) AddUserGRPC(user datastruct.User) (*endpoint.AddUserResponse, error){
	rq := &endpoint.AddUserRequest{
		Username:             user.Username,
	}
	addUserResponse, err := a.grpcClient.AddUser(a.ctx, rq)
	if err != nil {
		return nil, err
	}
	return addUserResponse, nil
}

func (a *Api) GetUserGRPC(userId int64) (*endpoint.GetUserResponse, error) {
	rq := &endpoint.GetUserRequest{
		Id:                   userId,
	}
	getUserResponse, err := a.grpcClient.GetUser(a.ctx, rq)
	if err != nil {
		return nil, err
	}
	return getUserResponse, nil
}

func (a *Api) CreateGarageGRPC() (*endpoint.CreateGarageResponse, error) {
	rq := &endpoint.EmptyRequest{}
	createGarageResponse, err := a.grpcClient.CreateGarage(a.ctx, rq)
	if err != nil {
		return nil, err
	}
	return createGarageResponse, nil
}

func (a *Api) GetGarageGRPC(garageId int64) (*endpoint.GetGarageResponse, error) {
	rq := &endpoint.GetGarageRequest{
		Id:                   garageId,
	}
	getGarageResponse, err := a.grpcClient.GetGarage(a.ctx, rq)
	if err != nil {
		return nil, err
	}
	return getGarageResponse, nil
}

func (a *Api) CarSearchGRPC(carModel datastruct.CarModel) (*endpoint.CarSearchResponse, error) {
	rq := &endpoint.CarSearchRequest{
		Brand:                carModel.Brand,
		Model:                carModel.Model,
	}
	carSearchResponse, err := a.grpcClient.CarSearch(a.ctx, rq)
	if err != nil {
		return nil, err
	}
	return carSearchResponse, nil
}

func (a *Api) AddToGarageGRPC(garageID int64, modelID int64) (*endpoint.AddToGarageResponse, error) {
	rq := &endpoint.AddToGarageRequest{
		GarageId:             garageID,
		ModelId:              modelID,
	}
	addToGarageResponse, err := a.grpcClient.AddToGarage(a.ctx, rq)
	if err != nil {
		return nil, err
	}
	return addToGarageResponse, nil
}