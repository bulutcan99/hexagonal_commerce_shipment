package psql

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/bulutcan99/commerce_shipment/internal/adapter/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

type DB struct {
	*pgxpool.Pool
	url          string
	QueryBuilder *squirrel.StatementBuilderType
}

func NewDB(ctx context.Context, Psql *config.PSQL) (*DB, error) {
	url := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable",
		Psql.Conn,
		Psql.User,
		Psql.Password,
		Psql.Host,
		Psql.Port,
		Psql.Name,
	)
	db, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}

	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	return &DB{
		db,
		url,
		&psql,
	}, nil
}

func (db *DB) Migrate() error {
	source, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		return err
	}
	fmt.Println("1", db.url)
	migration, err := migrate.NewWithSourceInstance("iofs", source, db.url)
	if err != nil {
		return err
	}
	fmt.Println("2")
	err = migration.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	fmt.Println("3")
	return nil
}

func (db *DB) Close() {
	db.Pool.Close()
}
