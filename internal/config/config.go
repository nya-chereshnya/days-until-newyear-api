package config

import (
	"log"
	"path/filepath"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"development"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	IP         string        `yaml:"IP" env-default:"localhost"`
	Port       string        `yaml:"port" env-default:"8080"`
	Timeout    time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimout time.Duration `yaml:"idle_timeout" env-default:"30s"`
}

func MustLoad() *Config {
	configPath, _ := filepath.Abs("../../internal/config/local.yaml")
	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("error reading config file: %s", err)
	}
	return &cfg
}
