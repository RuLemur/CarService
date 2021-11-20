package main

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/rulemur/CarService/internal/pkg"
)

func main() {
	app := pkg.NewApp()
	cfg := pkg.ReadConfig()
	app.RunApp(cfg)
}
