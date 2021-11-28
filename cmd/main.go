package main

import (
	"github.com/RuLemur/CarService/internal/logger"
	"github.com/RuLemur/CarService/internal/pkg"
	_ "github.com/jackc/pgx/stdlib"
	"os"
	"time"
)

var log = logger.NewDefaultLogger()

func main() {

	var app pkg.App

	var svc = pkg.ServiceKeeper{
		Services: []pkg.Service{
			&app.DB,
			&app.Rabbit,
		},
		ShutdownTimeout: time.Second * 10,
		PingPeriod:      time.Second * 30,
	}
	var application = pkg.Application{
		MainFunc:           app.RunApp,
		Resources:          &svc,
		TerminationTimeout: time.Second * 10,
	}
	if err := application.Run(); err != nil {
		log.Errorf("error: %s", err.Error())
		os.Exit(1)
	}
}
