package repository

import (
	"context"
	repository "github.com/bulutcan99/commerce_shipment/internal/adapter/storage/postgres"
	"github.com/bulutcan99/commerce_shipment/internal/core/domain"
	"time"
)

type PermissionRepository struct {
	db *repository.DB
}

func NewPermissionRepository(db *repository.DB) *PermissionRepository {
	return &PermissionRepository{
		db,
	}
}

func (p *PermissionRepository) Insert(ctx context.Context, permission *domain.Permission, userId uint64) (*domain.Permission, *domain.Error) {
	query := p.db.QueryBuilder.Insert("permissions").
		Values(permission.Entry, permission.AddFlag, permission.AdminFlag).
		Suffix("RETURNING *")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.ErrSqlInsert,
			Message: err.Error(),
		}
	}

	permission.CreatedAt = time.Now()
	permission.UpdatedAt = time.Now()
	err = p.db.QueryRow(ctx, sql, args...).Scan(
		&permission.ID,
		&permission.Entry,
		&permission.AddFlag,
		&permission.AdminFlag,
		&permission.CreatedAt,
		&permission.UpdatedAt,
	)

	if err != nil {
		return nil, &domain.Error{
			Code:    domain.ErrSqlInsert,
			Message: err.Error(),
		}
	}

	permissionUserQuery := p.db.QueryBuilder.Insert("user_permissions").
		Columns("user_id", "permission_id").
		Values(userId, permission.ID).
		Suffix("RETURNING *")
	sql, args, err = permissionUserQuery.ToSql()
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.ErrSqlInsert,
			Message: err.Error(),
		}
	}
	_, err = p.db.Exec(ctx, sql, args...)
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.ErrSqlInsert,
			Message: err.Error(),
		}
	}

	return permission, nil
}
