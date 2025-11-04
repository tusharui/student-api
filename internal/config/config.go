package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"addr"`
}

type Config struct {
	Env         string      `yaml:"env" env:"ENV"`
	StoragePath string      `yaml:"storage_path" env-required:"true"`
	HTTPServer  HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	// Try environment variable first
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flagConfig := flag.String("config", "", "path to the configuration file")
		flag.Parse()

		configPath = *flagConfig

		if configPath == "" {
			log.Fatal("config path is not set")
		}
	}

	// Check file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	// Read config using cleanenv
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("cannot read config file: %s", err)
	}

	return &cfg
}
