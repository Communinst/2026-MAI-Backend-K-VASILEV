package repository

import (
	"context"

	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/models"
)

type ProductRepositoryInterface interface {
	GetAll(ctx context.Context) ([]models.Product, error)
	Create(ctx context.Context, product *models.Product) (uint, error)
	GetByID(ctx context.Context, id uint64) (*models.Product, error)
	Update(ctx context.Context, id uint64, product *models.Product) error
	Delete(ctx context.Context, id uint64) error
	Search(ctx context.Context, query string) ([]models.Product, error)
}

type Repository struct {
	Product ProductRepositoryInterface
}
