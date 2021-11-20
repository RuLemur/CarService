package main

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/rulemur/car_service/internal/pkg"
)

func main() {
	app := pkg.NewApp()
	cfg := pkg.ReadConfig()
	app.RunApp(cfg)
}
