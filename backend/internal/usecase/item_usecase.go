package usecase

import (
	"context"

	"test_tablelink/internal/domain"
	"test_tablelink/internal/repository"
)

type ItemUsecase interface {
	Create(ctx context.Context, item *domain.Item) error
	Update(ctx context.Context, item *domain.Item) error
	GetAll(ctx context.Context, limit, offset int) ([]domain.Item, error)
	GetByUUID(ctx context.Context, uuid string) (*domain.Item, error)
	HardDelete(ctx context.Context, uuid string) error
}

type itemUsecase struct {
	itemRepo repository.ItemRepository
}

func NewItemUsecase(itemRepo repository.ItemRepository) ItemUsecase {
	return &itemUsecase{
		itemRepo: itemRepo,
	}
}

func (uc *itemUsecase) Create(ctx context.Context, item *domain.Item) error {
	return uc.itemRepo.Create(ctx, item)
}

func (uc *itemUsecase) Update(ctx context.Context, item *domain.Item) error {
	return uc.itemRepo.Update(ctx, item)
}

func (uc *itemUsecase) GetAll(ctx context.Context, limit, offset int) ([]domain.Item, error) {
	return uc.itemRepo.GetAll(ctx, limit, offset)
}

func (uc *itemUsecase) GetByUUID(ctx context.Context, uuid string) (*domain.Item, error) {
	return uc.itemRepo.GetByUUID(ctx, uuid)
}

func (uc *itemUsecase) HardDelete(ctx context.Context, uuid string) error {
	return uc.itemRepo.HardDelete(ctx, uuid)
}
