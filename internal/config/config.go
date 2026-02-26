package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer Config

type Config struct {
	Env         string `yaml:env env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:storage_path  env-required:"true" `
	HTTPServer  `yaml:http_server`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to config file")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			panic("config path is required")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("config file does not exist: %s", configPath)
	}

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatal("failed to read config file: %s", err)
	}

	return &cfg
}
