package repository

import (
	"context"
	repository "github.com/bulutcan99/commerce_shipment/internal/adapter/storage/postgres"
	"github.com/bulutcan99/commerce_shipment/internal/core/domain"
)

type PermissionRepository struct {
	db *repository.DB
}

func NewPermissionRepository(db *repository.DB) *PermissionRepository {
	return &PermissionRepository{
		db,
	}
}

func (p *PermissionRepository) AddPermission(ctx context.Context, permission *domain.Permission) (*domain.Permission, *domain.Error) {
	query := p.db.QueryBuilder.Insert("permissions").
		Columns("entry", "add_flag", "remove_flag", "admin_flag").
		Values(permission.Entry, permission.AddFlag, permission.RemoveFlag, permission.AdminFlag).
		Suffix("RETURNING *")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.ErrSqlInsert,
			Message: err.Error(),
		}
	}

	err = p.db.QueryRow(ctx, sql, args...).Scan(
		&permission.ID,
		&permission.Entry,
		&permission.AddFlag,
		&permission.RemoveFlag,
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

	return permission, nil
}
