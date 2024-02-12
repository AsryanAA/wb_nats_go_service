package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"wb_nats_go_service/internal/models"
)

const FilePath string = "configs/config.yaml"

type Config struct {
	models.DataBaseConfig `yaml:"db"`
	models.HTTPServer     `yaml:"server"`
}

// MustLoad функция считывания конфигурационного файла
func MustLoad() *Config {
	configPath := os.Getenv("FILE_PATH")
	if configPath == "" {
		// log.Fatal("CONFIG_PATH is no set")
		configPath = FilePath
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exist", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
