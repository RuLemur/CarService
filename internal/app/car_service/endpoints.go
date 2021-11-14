package car_service

import (
	"car-service/internal/app/car_service/endpoint"
	"google.golang.org/grpc"
)

type Module struct {
}

func NewModule() Module {
	return Module{}
}

// RunGRPC registers gRPC methods
func (m Module) RunGRPC(s *grpc.Server) {
	var srv *endpoint.GRPCRouter
	endpoint.RegisterCarServiceServer(s, srv)
}
