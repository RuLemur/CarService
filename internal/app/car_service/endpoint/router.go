package endpoint

import (
	"context"
)

type GRPCRouter struct {
}

func (G *GRPCRouter) GetGarageInfo(ctx context.Context, request *GetGarageInfoRequest) (*GetGarageInfoResponse, error) {
	garageInfo, err := GetGarageInfo(ctx, request.ID)
	if err != nil {
		return nil, err
	}
	return &GetGarageInfoResponse{
		Message: garageInfo.GarageName,
	}, nil
}
