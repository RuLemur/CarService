package pkg

import (
	"car-service"
	"google.golang.org/grpc"
)

// App ...
type App struct {
	Cfg             *CarService.Config
	GRPCServer      *grpc.Server
	GRPCInterceptor grpc.UnaryServerInterceptor
}

func (app *App) InitServer() {

}
