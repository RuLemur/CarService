package main

import (
	"car-service/internal/app/car_service"
	"flag"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"gopkg.in/yaml.v3"
	"log"
	"net"
	"os"
)

type Config struct {
	Server struct {
		GrpcHost string `yaml:"grpc_host"`
		Port     string `yaml:"port"`
		DBHost   string `yaml:"db"`
	}
}

func NewServerConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err = d.Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}

// ValidateConfigPath just makes sure, that the path provided is a file,
// that can be read
func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

// ParseFlags will create and parse the CLI flags
// and return the path to be used elsewhere
func ParseFlags() (string, error) {
	// String that contains the configured configuration path
	var configPath string

	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&configPath, "config", "./config/local.yml", "path to config file")

	// Actually parse the flags
	flag.Parse()

	// Validate the path first
	if err := ValidateConfigPath(configPath); err != nil {
		return "", err
	}

	// Return the configuration path
	return configPath, nil
}

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
	db, err := sqlx.Connect("pgx", config.Server.DBHost)
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

	//var g Garage
	//err = db.Get(&g, "SELECT * FROM garage LIMIT 1")
	//fmt.Println(g)

	listener, err := net.Listen("tcp", config.Server.Port)

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	service := car_service.NewModule()
	service.RunGRPC(grpcServer)

	grpcServer.Serve(listener)
}
