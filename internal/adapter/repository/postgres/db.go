package repository

import (
	"context"
	"fmt"
	"github.com/bulutcan99/shipment/internal/adapter/env"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	DB_USERNAME = &env.Env.DbUsername
	DB_PASSWORD = &env.Env.DbPassword
	DB_HOST     = &env.Env.Host
	DB_PORT     = &env.Env.DbPort
	DB_DATABASE = &env.Env.DbName
)

type DB struct {
	*pgxpool.Pool
}

func NewDB(ctx context.Context) (*DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		*DB_USERNAME,
		*DB_PASSWORD,
		*DB_HOST,
		*DB_PORT,
		*DB_DATABASE,
	)

	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return &DB{
		db,
	}, nil
}

func (db *DB) Close() {
	db.Pool.Close()
}
