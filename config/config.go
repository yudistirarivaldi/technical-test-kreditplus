package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server        ServerConfig
	DatabaseMysql DatabaseConfig
	JWT           JWTConfig
}

type JWTConfig struct {
	Secret string
}

type ServerConfig struct {
	Port            string
	ReadTimeout     int
	WriteTimeout    int
	ShutdownTimeout int
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		Server: ServerConfig{
			Port:            os.Getenv("SERVER_PORT"),
			ReadTimeout:     30,
			WriteTimeout:    30,
			ShutdownTimeout: 5,
		},
		DatabaseMysql: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASS"),
			Database: os.Getenv("DB_NAME"),
		},
		JWT: JWTConfig{
			Secret: os.Getenv("JWT_SECRET"),
		},
	}, nil
}
