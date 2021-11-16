package endpoint

import (
	"car-service/internal/app/datastruct"
	"context"
)

type GRPCRouter struct {
	s *Service
}

func NewGRPCRouter(s *Service) *GRPCRouter {
	return &GRPCRouter{s}
}

func (g *GRPCRouter) GetGarageInfo(ctx context.Context, request *GetGarageInfoRequest) (*GetGarageInfoResponse, error) {
	garageInfo, err := g.s.GetGarageInfo(ctx, request.ID)
	if err != nil {
		return nil, err
	}
	return &GetGarageInfoResponse{
		Message: garageInfo.GarageName,
	}, nil
}


func (g *GRPCRouter) AddUser(ctx context.Context, request *AddUserRequest) (*AddUserResponse, error) {
	userID, err := g.s.AddUser(ctx, datastruct.User{Username: request.Username})
	if err != nil {
		return nil, err
	}
	return &AddUserResponse{
		ID: userID,
	}, nil
}

