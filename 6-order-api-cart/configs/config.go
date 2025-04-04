package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB  DbConfig
	Jwt JwtConfig
}

type DbConfig struct {
	DSN string
}

type JwtConfig struct {
	Secret string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error while loading env: %s. Use default values.", err.Error())
		return &Config{
			DB: DbConfig{DSN: "def_val_1"},
			Jwt: JwtConfig{Secret: "def_val_2"},
		}
	}
	return &Config{
		DB: DbConfig{DSN: os.Getenv("DSN")},
		Jwt: JwtConfig{Secret: os.Getenv("SECRET")},
	}
}
