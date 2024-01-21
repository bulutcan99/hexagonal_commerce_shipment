package repository

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	repository "github.com/bulutcan99/commerce_shipment/internal/adapter/storage/postgres"
	"github.com/bulutcan99/commerce_shipment/internal/core/domain"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *repository.DB
}

func NewUserRepository(db *repository.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (u *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, *domain.Error) {
	query := psql.Insert("users").
		Columns("name", "surname", "email", "password", "address", "notification_radius", "role").
		Values(user.Name, user.Surname, user.Email, user.Password, user.Address, user.NotificationRadius, user.Role).
		Suffix("RETURNING *")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.ErrSqlInsert,
			Message: err.Error(),
		}
	}

	err = u.db.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Email,
		&user.Password,
		&user.Address,
		&user.NotificationRadius,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, &domain.Error{
			Code:    domain.ErrSqlInsert,
			Message: err.Error(),
		}
	}

	return user, nil
}

func (u *UserRepository) GetUserByID(ctx context.Context, id uint64) (*domain.User, *domain.Error) {
	var user domain.User

	query := psql.Select("*").
		From("users").
		Where(sq.Eq{"id": id}).
		Limit(1)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.ErrSqlSelect,
			Message: err.Error(),
		}
	}

	err = u.db.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Email,
		&user.Password,
		&user.Address,
		&user.NotificationRadius,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, &domain.Error{
				Code:    domain.DataNotFound,
				Message: err.Error(),
			}
		}
		return nil, &domain.Error{
			Code:    domain.ErrSqlSelect,
			Message: err.Error(),
		}
	}

	return &user, nil
}

func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, *domain.Error) {
	var user domain.User

	query := psql.Select("email", "password").
		From("users").
		Where(sq.Eq{"email": email}).
		Limit(1)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.ErrSqlSelect,
			Message: err.Error(),
		}
	}

	err = u.db.QueryRow(ctx, sql, args...).Scan(
		&user.Email,
		&user.Password,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, &domain.Error{
				Code:    domain.DataNotFound,
				Message: err.Error(),
			}
		}
		return nil, &domain.Error{
			Code:    domain.ErrSqlSelect,
			Message: err.Error(),
		}
	}

	return &user, nil
}

func (u *UserRepository) GetUsers(ctx context.Context) ([]domain.User, *domain.Error) {
	var user domain.User
	var users []domain.User

	query := psql.Select("*").
		From("users").
		OrderBy("id")

	sql, args, err := query.ToSql()
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, &domain.Error{
				Code:    domain.DataNotFound,
				Message: err.Error(),
			}
		}
		return nil, &domain.Error{
			Code:    domain.ErrSqlSelect,
			Message: err.Error(),
		}
	}

	rows, err := u.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.ErrSqlSelect,
			Message: err.Error(),
		}
	}

	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Surname,
			&user.Email,
			&user.Password,
			&user.Address,
			&user.NotificationRadius,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, &domain.Error{
				Code:    domain.ErrSqlSelect,
				Message: err.Error(),
			}
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserRepository) GetUsersWithLimit(ctx context.Context, skip, limit uint64) ([]domain.User, *domain.Error) {
	var user domain.User
	var users []domain.User

	query := psql.Select("*").
		From("users").
		OrderBy("id").
		Limit(limit).
		Offset(skip + 1)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.ErrSqlSelect,
			Message: err.Error(),
		}
	}

	rows, err := u.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.ErrSqlSelect,
			Message: err.Error(),
		}
	}

	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Surname,
			&user.Email,
			&user.Password,
			&user.Address,
			&user.NotificationRadius,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, &domain.Error{
				Code:    domain.ErrSqlSelect,
				Message: err.Error(),
			}
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *UserRepository) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, *domain.Error) {
	name := nullString(user.Name)
	surname := nullString(user.Surname)
	email := nullString(user.Email)
	password := nullString(user.Password)
	role := nullString(user.Role)
	address := nullString(user.Address)
	notificationRadius := nullUint64(user.NotificationRadius)

	query := psql.Update("users").
		Set("name", sq.Expr("COALESCE(?, name)", name)).
		Set("surname", sq.Expr("COALESCE(?, surname)", surname)).
		Set("email", sq.Expr("COALESCE(?, email)", email)).
		Set("role", sq.Expr("COALESCE(?, role)", role)).
		Set("password", sq.Expr("COALESCE(?, password)", password)).
		Set("address", sq.Expr("COALESCE(?, address)", address)).
		Set("notification_radius", sq.Expr("COALESCE(?, notification_radius)", notificationRadius)).
		Set("updated_at", sq.Expr("NOW()")).
		Where(sq.Eq{"id": user.ID}).
		Suffix("RETURNING *")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.ErrSqlUpdate,
			Message: err.Error(),
		}
	}
	err = u.db.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Email,
		&user.Password,
		&user.Address,
		&user.NotificationRadius,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, &domain.Error{
			Code:    domain.ErrSqlUpdate,
			Message: err.Error(),
		}
	}

	return user, nil
}

func (u *UserRepository) DeleteUser(ctx context.Context, id uint64) *domain.Error {
	query := psql.Delete("users").
		Where(sq.Eq{"id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return &domain.Error{
			Code:    domain.ErrSqlDelete,
			Message: err.Error(),
		}
	}

	_, err = u.db.Exec(ctx, sql, args...)
	if err != nil {
		return &domain.Error{
			Code:    domain.ErrSqlDelete,
			Message: err.Error(),
		}
	}

	return nil
}