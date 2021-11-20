package main

import (
	"CarService/internal/pkg"
	_ "github.com/jackc/pgx/stdlib"
)

func main() {
	app := pkg.NewApp()
	cfg := pkg.ReadConfig()
	app.RunApp(cfg)
}
