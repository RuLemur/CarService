package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"sync"
)

type Config interface {
	GetConfig() *ServerConfig
}

type config struct {
	cfg *ServerConfig
}

var (
	instance *config
	mu       = new(sync.Mutex)
)

func GetInstance() Config {
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		cfg := readConfig()
		instance = &config{cfg: cfg}
	}
	return instance
}

type ServerConfig struct {
	Server struct {
		Host string `yaml:"host"`
	} `yaml:"server"`

	Database struct {
		DBHost string `yaml:"db"`
	} `yaml:"database"`

	Queue struct {
		Host string `yaml:"host"`
	} `yaml:"queue"`

	Middleware struct {
		Host string `yaml:"host"`
	} `yaml:"middleware"`
}

func newServerConfig(configPath string) (*ServerConfig, error) {
	serverConfig := &ServerConfig{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err = d.Decode(&serverConfig); err != nil {
		return nil, err
	}
	return serverConfig, nil
}

// validateConfigPath just makes sure, that the path provided is a file,
// that can be read
func validateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

// parseFlags will create and parse the CLI flags
// and return the path to be used elsewhere
func parseFlags() (string, error) {
	// String that contains the configured configuration path
	var configPath string

	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&configPath, "config", "./config/local.yml", "path to config file")

	// Actually parse the flags
	flag.Parse()

	// Validate the path first
	if err := validateConfigPath(configPath); err != nil {
		return "", err
	}

	// Return the configuration path
	return configPath, nil
}

func readConfig() *ServerConfig {
	cfgPath, err := parseFlags()

	if err != nil {
		log.Fatal(err)
	}
	cfg, err := newServerConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}

func (s *config) GetConfig() *ServerConfig {
	return s.cfg
}
