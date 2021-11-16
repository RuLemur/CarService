package car_service

import (
	"car-service/internal/app/car_service/endpoint"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

type Module struct {
}

func NewModule() Module {
	return Module{}
}

// RunGRPC registers gRPC methods
func (m Module) RunGRPC(db *sqlx.DB, s *grpc.Server) {
	service := endpoint.NewService(db)
	srv := endpoint.NewGRPCRouter(service)
	endpoint.RegisterCarServiceServer(s, srv)
}
