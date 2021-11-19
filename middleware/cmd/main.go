package main

import (
	config "car-service/internal/pkg"
	"car-service/middleware/internal/pkg"
)

//import

func main() {
	app := pkg.NewApp()
	cfg := config.ReadConfig()
	app.RunApp(cfg)
}