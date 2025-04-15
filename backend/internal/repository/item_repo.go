package repository

import (
	"context"
	"fmt"
	"time"

	"test_tablelink/internal/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ItemRepository interface {
	Create(ctx context.Context, item *domain.Item) error
	Update(ctx context.Context, item *domain.Item) error
	GetAll(ctx context.Context, limit, offset int) ([]domain.Item, error)
	GetByUUID(ctx context.Context, uuid string) (*domain.Item, error)
	HardDelete(ctx context.Context, uuid string) error
}

type ItemRepo struct {
	db *pgxpool.Pool
}

func NewItemRepo(db *pgxpool.Pool) *ItemRepo {
	return &ItemRepo{db: db}
}

func (r *ItemRepo) Create(ctx context.Context, item *domain.Item) error {
	item.UUID = uuid.NewString()
	now := time.Now()
	item.CreatedAt = now
	item.UpdatedAt = now

	query := `
		INSERT INTO tm_item (
			uuid, name, price, status,
			created_at, updated_at, deleted_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := r.db.Exec(ctx, query,
		item.UUID,
		item.Name,
		item.Price,
		item.Status,
		item.CreatedAt,
		item.UpdatedAt,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to insert item: %w", err)
	}
	return nil
}

func (r *ItemRepo) Update(ctx context.Context, item *domain.Item) error {
	now := time.Now()
	item.UpdatedAt = now

	query := `
		UPDATE tm_item
		SET name = $1, price = $3, status = $4, updated_at = $5
		WHERE uuid = $6 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(ctx, query,
		item.Name,
		item.Price,
		item.Status,
		item.UpdatedAt,
		item.UUID,
	)
	if err != nil {
		return fmt.Errorf("failed to update item: %w", err)
	}
	return nil
}

func (r *ItemRepo) GetAll(ctx context.Context, limit, offset int) ([]domain.Item, error) {
	query := `
		SELECT uuid, name, price, status,
		       created_at, updated_at, deleted_at
		FROM tm_item
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to select items: %w", err)
	}
	defer rows.Close()

	var items []domain.Item
	for rows.Next() {
		var item domain.Item
		if err := rows.Scan(
			&item.UUID,
			&item.Name,
			&item.Price,
			&item.Status,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.DeletedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan item: %w", err)
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *ItemRepo) GetByUUID(ctx context.Context, uuid string) (*domain.Item, error) {
	query := `
		SELECT uuid, name, price, status,
		       created_at, updated_at, deleted_at
		FROM tm_item
		WHERE uuid = $1 AND deleted_at IS NULL
	`
	var item domain.Item
	err := r.db.QueryRow(ctx, query, uuid).Scan(
		&item.UUID,
		&item.Name,
		&item.Price,
		&item.Status,
		&item.CreatedAt,
		&item.UpdatedAt,
		&item.DeletedAt,
	)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, nil // No item found
		}
		return nil, fmt.Errorf("failed to query item by UUID: %w", err)
	}
	return &item, nil
}

func (r *ItemRepo) HardDelete(ctx context.Context, uuid string) error {
	query := `
		UPDATE tm_item
		SET deleted_at = NOW()
		WHERE uuid = $1
	`
	_, err := r.db.Exec(ctx, query, uuid)
	if err != nil {
		return fmt.Errorf("failed to hard delete item: %w", err)
	}
	return nil
}
