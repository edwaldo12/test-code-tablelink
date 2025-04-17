package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"test_tablelink/internal/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type IngredientRepository interface {
	Create(ctx context.Context, ing *domain.Ingredient) error
	Update(ctx context.Context, ing *domain.Ingredient) error
	GetAll(ctx context.Context, limit, offset int) ([]domain.Ingredient, error)
	GetByName(ctx context.Context, name string) (*domain.Ingredient, error)
	HardDelete(ctx context.Context, uuid string) error
}

type IngredientRepo struct {
	db *pgxpool.Pool
}

func NewIngredientRepo(db *pgxpool.Pool) *IngredientRepo {
	return &IngredientRepo{db: db}
}

func (r *IngredientRepo) Create(ctx context.Context, ing *domain.Ingredient) error {
	ing.UUID = uuid.NewString()
	now := time.Now()
	ing.CreatedAt = sql.NullTime{Time: now, Valid: true}
	ing.UpdatedAt = sql.NullTime{Time: now, Valid: true}

	query := `
			INSERT INTO tm_ingredient (
					uuid, name, cause_alergy, type, status,
					created_at, updated_at, deleted_at
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := r.db.Exec(ctx, query,
		ing.UUID,
		ing.Name,
		ing.CauseAllergy,
		ing.Type,
		ing.Status,
		ing.CreatedAt,
		ing.UpdatedAt,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to insert ingredient: %w", err)
	}
	return nil
}

func (r *IngredientRepo) GetAll(ctx context.Context, limit, offset int) ([]domain.Ingredient, error) {
	query := `
			SELECT uuid, name, cause_alergy, type, status,
						 created_at, updated_at, deleted_at
			FROM tm_ingredient
			WHERE deleted_at IS NULL
			ORDER BY created_at DESC
			LIMIT $1
			OFFSET $2
	`
	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to select ingredients: %w", err)
	}
	defer rows.Close()

	var ingredients []domain.Ingredient
	for rows.Next() {
		var ing domain.Ingredient
		if err := rows.Scan(
			&ing.UUID,
			&ing.Name,
			&ing.CauseAllergy,
			&ing.Type,
			&ing.Status,
			&ing.CreatedAt,
			&ing.UpdatedAt,
			&ing.DeletedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan ingredient: %w", err)
		}
		ingredients = append(ingredients, ing)
	}
	return ingredients, nil
}

func (r *IngredientRepo) HardDelete(ctx context.Context, uuid string) error {
	query := `
			UPDATE tm_ingredient
			SET deleted_at = NOW()
			WHERE uuid = $1
	`
	_, err := r.db.Exec(ctx, query, uuid)
	if err != nil {
		return fmt.Errorf("failed to hard delete ingredient: %w", err)
	}
	return nil
}

func (r *IngredientRepo) GetByName(ctx context.Context, name string) (*domain.Ingredient, error) {
	query := `
			SELECT uuid, name, cause_alergy, type, status,
						 created_at, updated_at, deleted_at
			FROM tm_ingredient
			WHERE name = $1 AND deleted_at IS NULL
	`
	row := r.db.QueryRow(ctx, query, name)

	var ing domain.Ingredient
	if err := row.Scan(
		&ing.UUID,
		&ing.Name,
		&ing.CauseAllergy,
		&ing.Type,
		&ing.Status,
		&ing.CreatedAt,
		&ing.UpdatedAt,
		&ing.DeletedAt,
	); err != nil {
		return nil, fmt.Errorf("failed to scan ingredient: %w", err)
	}
	return &ing, nil
}

func (r *IngredientRepo) Update(ctx context.Context, ing *domain.Ingredient) error {
	ing.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	query := `
			UPDATE tm_ingredient
			SET name = $1, cause_alergy = $2, type = $3, status = $4,
					updated_at = $5
			WHERE uuid = $6
	`
	_, err := r.db.Exec(ctx, query,
		ing.Name,
		ing.CauseAllergy,
		ing.Type,
		ing.Status,
		ing.UpdatedAt,
		ing.UUID,
	)
	if err != nil {
		return fmt.Errorf("failed to update ingredient: %w", err)
	}
	return nil
}
