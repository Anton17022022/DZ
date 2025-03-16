package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db             Db
}

type Db struct {
	Db string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(err.Error())
	} 
	envParams := LoadEnv()
	return &Config{
		Db: Db{
			Db: envParams["DSN_PG_Orders"],
		},
	}
}

func LoadEnv() map[string]string {
	envParams := make(map[string]string, 0)
	envs := []string{"DSN_PG_Orders"}
	for _, value := range envs{
		env, ok := os.LookupEnv(value)
		if ok {
			envParams[value] = env
		} else {
			envParams[value] = ""
		}
	}
	return envParams
}
