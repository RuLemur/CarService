package main

import (
	"github.com/RuLemur/CarService/internal/logger"
	"github.com/RuLemur/CarService/internal/pkg"
	_ "github.com/jackc/pgx/stdlib"
	"os"
	"time"
)

var (
	log = logger.NewDefaultLogger()
)

const (
	shutdownTimeout    = time.Second * 10
	pingPeriod         = time.Minute
	terminationTimeout = time.Second * 10
)

func main() {

	var app pkg.App

	log.Infof("Starting service...")
	var svc = pkg.ServiceKeeper{
		Services: []pkg.Service{
			&app.DB,
			&app.Rabbit,
		},
		ShutdownTimeout: shutdownTimeout,
		PingPeriod:      pingPeriod,
	}
	var application = pkg.Application{
		MainFunc:           app.RunApp,
		Resources:          &svc,
		TerminationTimeout: terminationTimeout,
	}
	if err := application.Run(); err != nil {
		log.Errorf("error: %s", err.Error())
		os.Exit(1)
	}
}
