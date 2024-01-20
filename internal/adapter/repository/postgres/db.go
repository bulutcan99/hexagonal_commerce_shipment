package repository

import (
	"context"
	"fmt"
	"github.com/bulutcan99/commerce_shipment/internal/adapter/env"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	DbUsername = &env.Env.DbUsername
	DbPass     = &env.Env.DbPassword
	DbHost     = &env.Env.Host
	DbPort     = &env.Env.DbPort
	Database   = &env.Env.DbName
)

type DB struct {
	*pgxpool.Pool
}

func NewDB(ctx context.Context) (*DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		*DbUsername,
		*DbPass,
		*DbHost,
		*DbPort,
		*Database,
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
