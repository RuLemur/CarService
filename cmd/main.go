package main

import (
	"car_service/internal/pkg"
	_ "github.com/jackc/pgx/stdlib"
)

func main() {
	app := pkg.NewApp()
	cfg := pkg.ReadConfig()
	app.RunApp(cfg)
}
