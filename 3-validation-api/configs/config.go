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
	Host     string
}

type StatusCodes struct {
	StatusCodeBadRequest string
	StatusCodeOk         string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error while loading .env file, using default config")
	}
	envParams := LoadEnv()
	return &Config{
		Verify: VerifyConfig{
			Email:    envParams["Email"],
			Password: envParams["Pswd"],
			Host:     envParams["Host"],
		},
		StatusResponce: StatusCodes{
			StatusCodeBadRequest: envParams["StatusCodeBadRequest"],
			StatusCodeOk:         envParams["StatusCodeOk"],
		},
	}
}

func LoadEnv() map[string]string {
	envParams := make(map[string]string, 0)
	envs := []string{"Email", "Pswd", "Host", "StatusCodeOk", "StatusCodeBadRequest"}
	for _, value := range envs {
		param, getResult := os.LookupEnv(value)
		if !getResult {
			log.Printf("%s can't get from env\n", param)
			envParams[param] = ""
		} else {
			envParams[param] = value
		}
	}
	return envParams
}
