package pkg

import (
	"context"
	"github.com/RuLemur/CarService/internal/app/car_service"
	"github.com/RuLemur/CarService/internal/app/car_service/router"
	"github.com/RuLemur/CarService/internal/config"
	"github.com/RuLemur/CarService/internal/logger"
	"github.com/RuLemur/CarService/internal/queue"
	"github.com/RuLemur/CarService/internal/repo"
	"github.com/RuLemur/CarService/pkg/endpoint"
	"google.golang.org/grpc"
	"net"
	"time"
)

var log = logger.NewDefaultLogger()

type App struct {
	Rabbit queue.Client
	DB     repo.Repository
}

func (app *App) RunApp(ctx context.Context, halt <-chan struct{}) error {
	cfg := config.GetInstance().GetConfig()

	log.Infof("Connect to tcp with: %s", cfg.Server.Host)
	listener, err := net.Listen("tcp", cfg.Server.Host)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(
		app.WithServerUnaryInterceptor(),
	)

	srv := car_service.NewService(
		app.DB,
		app.Rabbit)
	rtr := router.NewGRPCRouter(srv)
	endpoint.RegisterCarServiceServer(grpcServer, rtr)

	var errShutdown = make(chan error, 1)
	go func() {
		defer close(errShutdown)
		select {
		case <-halt:
		case <-ctx.Done():
		}
		grpcServer.Stop()
		log.Infof("Down GRPC server")
	}()
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	err, ok := <-errShutdown
	if ok {
		return err
	}
	return nil
}

func (app *App) WithServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(app.ServerInterceptor)
}

func (app *App) ServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	log.Infof("Request - Method: %s\tDuration:%s\t",
		info.FullMethod,
		time.Since(start))
	log.Infof("Request: %s\t", req)

	h, err := handler(ctx, req)

	log.Infof("Response: %s\tError:%v\n", h, err)

	return h, err
}
