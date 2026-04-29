package repository

import (
	"context"

	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/models"
)

type ProductRepositoryInterface interface {
	GetAll(ctx context.Context) ([]models.Product, error)
	Create(ctx context.Context, product *models.Product) (uint, error)
	GetByID(ctx context.Context, id uint64) (*models.Product, error)
	Search(ctx context.Context, query string) ([]models.Product, error)
}

type Repository struct {
	Product ProductRepositoryInterface
}
