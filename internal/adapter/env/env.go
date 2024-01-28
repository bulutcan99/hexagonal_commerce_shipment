package env

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"os"
	"sync"
)

type ENV struct {
	Host              string `env:"HOST,required"`
	ServerPort        int    `env:"SERVER_PORT,required"`
	ServerReadTimeout int    `env:"SERVER_READ_TIMEOUT,required"`
	TokenSymmetric    string `env:"TOKEN_SYMMETRIC,required"`
	TokenTTL          string `env:"TOKEN_TTL,required"`
	DbConn            string `env:"DB_CONN,required"`
	DbPort            int    `env:"DB_PORT,required"`
	DbUsername        string `env:"DB_USERNAME,required"`
	DbPassword        string `env:"DB_PASSWORD,required"`
	DbName            string `env:"DB_NAME,required"`
	RedisPort         int    `env:"REDIS_PORT,required"`
	RedisPassword     string `env:"REDIS_PASSWORD,required"`
	RedisDBNumber     int    `env:"REDIS_DB_NUMBER,required"`
}

var doOnce sync.Once
var Env ENV

func ParseEnv() *ENV {
	doOnce.Do(func() {
		err := godotenv.Load()
		if err != nil {
			fmt.Printf("error loading .env file: %v", err)
			os.Exit(0)
		}
		if err := env.Parse(&Env); err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(0)
		}
	})
	return &Env
}
