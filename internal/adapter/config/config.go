package config

import (
	"github.com/bulutcan99/commerce_shipment/internal/adapter/env"
)

var (
	Host           = &env.Env.Host
	ServerPort     = &env.Env.ServerPort
	TokenSymmetric = &env.Env.TokenSymmetric
	TokenTTL       = &env.Env.TokenTTL
	DbPort         = &env.Env.DbPort
	DbUsername     = &env.Env.DbUsername
	DbPassword     = &env.Env.DbPassword
	DbName         = &env.Env.DbName
	DbConn         = &env.Env.DbConn
	RedisPort      = &env.Env.RedisPort
	RedisPassword  = &env.Env.RedisPassword
	RedisDBNumber  = &env.Env.RedisDBNumber
	LogLevel       = &env.Env.LogLevel
)

type (
	Container struct {
		App   *App
		Token *Token
		Redis *Redis
		PSQL  *PSQL
		Fiber *Fiber
		Kafka *Kafka
	}

	App struct {
		Name string
	}

	Token struct {
		SymmetricKey string
		TTL          string
	}

	Redis struct {
		Host     string
		Port     int
		Password string
		DbNumber int
	}

	PSQL struct {
		Conn     string
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	}

	Fiber struct {
		Host string
		Port int
	}

	Kafka struct {
		Host  string
		Port  int
		Topic string
	}
)

func New() *Container {
	app := &App{
		Name: "commerce_shipment",
	}

	token := &Token{
		SymmetricKey: *TokenSymmetric,
		TTL:          *TokenTTL,
	}

	redis := &Redis{
		Host:     *Host,
		Port:     *RedisPort,
		Password: *RedisPassword,
		DbNumber: *RedisDBNumber,
	}

	psql := &PSQL{
		Conn:     *DbConn,
		Host:     *Host,
		Port:     *DbPort,
		User:     *DbUsername,
		Password: *DbPassword,
		Name:     *DbName,
	}

	fiber := &Fiber{
		Host: *Host,
		Port: *ServerPort,
	}

	return &Container{
		App:   app,
		Token: token,
		Redis: redis,
		PSQL:  psql,
		Fiber: fiber,
	}
}
