package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	repository "github.com/bulutcan99/commerce_shipment/internal/adapter/storage/postgres"
	"github.com/bulutcan99/commerce_shipment/internal/core/domain"
)

type CategoryRepository struct {
	db *repository.DB
}

func NewCategoryRepository(db *repository.DB) *CategoryRepository {
	return &CategoryRepository{
		db,
	}
}

func (c *CategoryRepository) CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, *domain.Error) {
	query := psql.Insert("categories").
		Columns("name").
		Values(category.Name).
		Suffix("RETURNING *")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.ErrSqlInsert,
			Message: err.Error(),
		}
	}

	err = c.db.QueryRow(ctx, sql, args...).Scan(
		&category.ID,
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		return nil, &domain.Error{
			Code:    domain.ErrSqlInsert,
			Message: err.Error(),
		}
	}
	return category, nil
}

func (c *CategoryRepository) GetCategoryByID(ctx context.Context, id uint64) (*domain.Category, *domain.Error) {
	var category domain.Category

	query := psql.Select("*").
		From("categories").
		Where(sq.Eq{"id": id}).
		Limit(1)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.ErrSqlSelect,
			Message: err.Error(),
		}
	}

	err = c.db.QueryRow(ctx, sql, args...).Scan(
		&category.ID,
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		return nil, &domain.Error{
			Code:    domain.ErrSqlSelect,
			Message: err.Error(),
		}
	}

	return &category, nil
}

func (c *CategoryRepository) GetCategories(ctx context.Context) ([]*domain.Category, *domain.Error) {
	var categories []*domain.Category

	query := psql.Select("*").
		From("categories")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.ErrSqlSelect,
			Message: err.Error(),
		}
	}

	rows, err := c.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.ErrSqlSelect,
			Message: err.Error(),
		}
	}

	for rows.Next() {
		var category domain.Category
		err = rows.Scan(
			&category.ID,
			&category.Name,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			return nil, &domain.Error{
				Code:    domain.ErrSqlSelect,
				Message: err.Error(),
			}
		}
		categories = append(categories, &category)
	}

	return categories, nil
}

func (c *CategoryRepository) UpdateCategory(ctx context.Context, category *domain.Category) (*domain.Category, *domain.Error) {
	query := psql.Update("categories").
		Set("name", category.Name).
		Where(sq.Eq{"id": category.ID}).
		Suffix("RETURNING *")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, &domain.Error{
			Code:    domain.ErrSqlUpdate,
			Message: err.Error(),
		}
	}

	err = c.db.QueryRow(ctx, sql, args...).Scan(
		&category.ID,
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		return nil, &domain.Error{
			Code:    domain.ErrSqlUpdate,
			Message: err.Error(),
		}
	}

	return category, nil
}

func (c *CategoryRepository) DeleteCategory(ctx context.Context, id uint64) *domain.Error {
	query := psql.Delete("categories").
		Where(sq.Eq{"id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return &domain.Error{
			Code:    domain.ErrSqlDelete,
			Message: err.Error(),
		}
	}

	_, err = c.db.Exec(ctx, sql, args...)
	if err != nil {
		return &domain.Error{
			Code:    domain.ErrSqlDelete,
			Message: err.Error(),
		}
	}

	return nil
}
