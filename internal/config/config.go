package config

import (
	// "log"
	"os"
)

type Config struct {
	Server struct {
		Port string
	}
	Database struct {
		DSN string
	}
}

func Load() *Config {
	cfg := &Config{}
	cfg.Server.Port = os.Getenv("SERVER_PORT")
	if cfg.Server.Port == "" {
		cfg.Server.Port = "8080"
	}
	cfg.Database.DSN = os.Getenv("DATABASE_DSN")
	// if cfg.Database.DSN == "" {
	// 	log.Fatal("DATABASE_DSN not set")
	// }
	return cfg
}
