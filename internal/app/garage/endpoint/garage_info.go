package endpoint

import (
	garage "car-service/internal/app/garage/datastruct"
	"context"
	"fmt"
)

func GetGarageInfo(ctx context.Context, garageID int64) (garage.GarageInfo, error) {
	return garage.GarageInfo{GarageName: fmt.Sprintf("Hello!, %d", garageID)}, nil
}
