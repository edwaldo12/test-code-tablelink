package usecase

import (
	"context"
	"fmt"
	"test_tablelink/internal/domain"
	"test_tablelink/internal/repository"
)

type IngredientUsecase interface {
	Create(ctx context.Context, ing *domain.Ingredient) error
	Update(ctx context.Context, ing *domain.Ingredient) error
	GetAll(ctx context.Context, limit, offset int) ([]domain.Ingredient, error)
	HardDelete(ctx context.Context, uuid string) error
}

type ingredientUsecase struct {
	ingredientRepo repository.IngredientRepository
}

func NewIngredientUsecase(ingredientRepo repository.IngredientRepository) IngredientUsecase {
	return &ingredientUsecase{
		ingredientRepo: ingredientRepo,
	}
}

func (uc *ingredientUsecase) Create(ctx context.Context, ing *domain.Ingredient) error {
	existing, err := uc.ingredientRepo.GetByName(ctx, ing.Name)
	if err != nil {
		return err
	}
	if existing != nil {
		return fmt.Errorf("ingredient with name %s already exists", ing.Name)
	}

	err = uc.ingredientRepo.Create(ctx, ing)
	if err != nil {
		return fmt.Errorf("failed to create ingredient: %w", err)
	}
	return nil
}

func (uc *ingredientUsecase) Update(ctx context.Context, ing *domain.Ingredient) error {
	existing, _ := uc.ingredientRepo.GetByName(ctx, ing.Name)
	if existing != nil && existing.UUID != ing.UUID {
		return fmt.Errorf("name already exists")
	}
	err := uc.ingredientRepo.Update(ctx, ing)
	if err != nil {
		return fmt.Errorf("failed to update ingredient: %w", err)
	}
	return nil
}

func (uc *ingredientUsecase) GetAll(ctx context.Context, limit, offset int) ([]domain.Ingredient, error) {
	ingredients, err := uc.ingredientRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch ingredients: %w", err)
	}
	return ingredients, nil
}

func (uc *ingredientUsecase) HardDelete(ctx context.Context, uuid string) error {
	err := uc.ingredientRepo.HardDelete(ctx, uuid)
	if err != nil {
		return fmt.Errorf("failed to hard delete ingredient: %w", err)
	}
	return nil
}
