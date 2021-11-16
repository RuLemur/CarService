package endpoint

import (
	"car-service/internal/app/datastruct"
	"car-service/internal/repo"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *Service {
	return &Service{db}
}

func (s *Service) GetGarageInfo(ctx context.Context, garageID int64) (datastruct.GarageInfo, error) {
	return datastruct.GarageInfo{GarageName: fmt.Sprintf("Hello!, %d", garageID)}, nil
}

func (s *Service) AddUser(ctx context.Context, user datastruct.User) (int64, error) {
	err := repo.AddNewUser(s.db, &user)
	return user.ID, err
}
