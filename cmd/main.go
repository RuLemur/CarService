package main

import (
	"car-service/internal/app/garage"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"net"
)

func main() {
	cfgPath, err := ParseFlags()
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := NewServerConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	// Run the server
	cfg.Run()
}

func (config Config) Run() {

	// connect to the database
	db, err := sqlx.Connect("pgx", config.db)
	if err != nil {
		log.Fatalln(err)
	}

	//db.QueryLogFunc = logDBQuery(logger)
	//db.ExecLogFunc = logDBExec(logger)
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	type Garage struct {
		Id         int    `db:"id"`
		GarageName string `db:"garage_name"`
	}

	var g Garage
	err = db.Get(&g, "SELECT * FROM garage LIMIT 1")
	fmt.Println(g)

	listener, err := net.Listen("tcp", config.Port)

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	service := garage.NewModule()
	service.RunGRPC(grpcServer)

	grpcServer.Serve(listener)
}
