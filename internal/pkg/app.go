package pkg

import (
	"car-service/internal/app/car_service"
	"car-service/internal/app/car_service/endpoint"
	"car-service/internal/queue"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"net"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (app App) RunApp(config *Config) {
	// connect to the database
	db, err := sqlx.Connect("pgx", config.Database.DBHost)

	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	//connect to Rabbit MQ
	rabbit := queue.NewClient(config.Queue.Host)
	err = rabbit.ConnectToServer()
	if err != nil {
		log.Fatalln(err)
	}
	defer rabbit.CloseConnect()

	// grpc Server
	listener, err := net.Listen("tcp",  config.Server.Host)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	service := car_service.NewService(db, rabbit)
	srv := endpoint.NewGRPCRouter(service)
	endpoint.RegisterCarServiceServer(grpcServer, srv)

	err = grpcServer.Serve(listener)
	if err != nil {
		panic("failed to Serve server")
	}
}
