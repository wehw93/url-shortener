package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env string `yaml:"env" env-required:"true"`

	StorageConfig `yaml:"storage"`
	HTTPServer    `yaml:"http_server"`
}

type HTTPServer struct {
	Addres       string        `yaml:"address" env-default:"localhost:8080"`
	Timeout      time.Duration `yaml:"timeout" env-default:"4s"`
	Idle_Timeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	User         string        `yaml:"user" env-required:"true"`
	Password     string        `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

type StorageConfig struct {
	Type string `yaml:"type" env:"STORAGE_TYPE" env-default:"sqlite"`

	SQLite struct {
		Path string `yaml:"path" env-default:"./storage.db"`
	} `yaml:"sqlite"`

	MongoDB struct {
		URI string `yaml:"uri" env-default:"mongodb://localhost:27017"`
	} `yaml:"mongodb"`
}

func MustLoad() *Config {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatal("error of Load .env")
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s is not exist", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config file %s ", err)
	}
	return &cfg
}
