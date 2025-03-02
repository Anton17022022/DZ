package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Verify         VerifyConfig
	StatusResponce StatusCodes
}

type VerifyConfig struct {
	Email    string
	Password string
	Host string
}

type StatusCodes struct {
	StatusCodeBadRequest string
	StatusCodeOk         string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error while loading .env file, using default config")
	}
	return &Config{
		Verify: VerifyConfig{
			Email:    os.Getenv("Email"),
			Password: os.Getenv("Pswd"),
			Host:     os.Getenv("Host"),
		},
		StatusResponce: StatusCodes{
			StatusCodeBadRequest: os.Getenv("StatusCodeBadRequest"),
			StatusCodeOk:         os.Getenv("StatusCodeOk"),
		},
	}
}
