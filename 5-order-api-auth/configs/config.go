package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Jwt JwtSecret
	Db  DbConfig
}

type JwtSecret struct {
	Secret string
}

type DbConfig struct {
	Dsn string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error %s. Using default value.", err.Error())
		return &Config{
			Jwt: JwtSecret{"def_value_1"},
			Db: DbConfig{"def_value_2"},
		}
	}
	return &Config{
		Jwt: JwtSecret{os.Getenv("SECRET")},
		Db: DbConfig{os.Getenv("DSN")},
	}
}
