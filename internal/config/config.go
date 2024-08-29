package config

import (
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"os"
	_ "os"
)

type Config struct {
	DatabaseUrl   string
	ServerAddress string
	LogLevel      string
	SecretKey     string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	config := &Config{
		DatabaseUrl:   os.Getenv("DATABASE_URL"),
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
		LogLevel:      os.Getenv("LOG_LEVEL"),
		SecretKey:     os.Getenv("SECRET_KEY"),
	}

	return config, nil

}
