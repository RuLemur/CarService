package pkg

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/rulemur/CarService/internal/app/car_service"
	"github.com/rulemur/CarService/internal/app/car_service/endpoint"
	"github.com/rulemur/CarService/internal/queue"
	"github.com/rulemur/CarService/internal/repo"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"time"
)

type App struct {
	log *logrus.Logger
}

func NewApp() *App {
	return &App{}
}

func (app *App) RunApp(config *Config) {
	app.log = logrus.New()
	app.initLogger(app.log)

	// connect to the database
	app.log.Println("Connecting to database...")
	dbConnection, err := sqlx.Connect("pgx", config.Database.DBHost)
	if err != nil {
		app.log.Fatalln(err)
	}
	db := &repo.QueryLogger{Queryer: dbConnection, Logger: app.log}
	app.log.Println("Connected.")
	defer dbConnection.Close()

	//connect to Rabbit MQ
	app.log.Println("Connecting to RabbitMQ Server...")
	rabbit := queue.NewClient(config.Queue.Host)
	err = rabbit.ConnectToServer()
	if err != nil {
		app.log.Fatalln(err)
	}
	app.log.Println("Connected.")
	defer rabbit.CloseConnect()

	// grpc Server
	app.log.Println("Starting GRPC Server...")
	listener, err := net.Listen("tcp", config.Server.Host)
	if err != nil {
		app.log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		app.withServerUnaryInterceptor(),
	)

	service := car_service.NewService(db, rabbit)
	srv := endpoint.NewGRPCRouter(service)
	endpoint.RegisterCarServiceServer(grpcServer, srv)

	app.log.Println("Started. Listen requests")
	err = grpcServer.Serve(listener)
	if err != nil {
		panic("failed to Serve server")
	}
}

func (app *App) initLogger(log *logrus.Logger) {
	log.SetFormatter(&logrus.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
}

func (app *App) withServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(app.serverInterceptor)
}

func (app *App) serverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	app.log.Infof("Request - Method: %s\tDuration:%s\t",
		info.FullMethod,
		time.Since(start))
	app.log.Infof("Request: %s\t", req)

	// Calls the handler
	h, err := handler(ctx, req)

	app.log.Infof("Response: %s\tError:%v\n", h, err)

	return h, err
}
