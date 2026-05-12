package service

import (
	"context"

	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/config"
	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/models"
	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/repository"
)

type ProductServiceInterface interface {
	GetProducts(ctx context.Context) ([]models.Product, error)
	CreateProduct(ctx context.Context, product *models.Product) (uint, error)
	GetProductByID(ctx context.Context, id uint64) (*models.Product, error)
	UpdateProduct(ctx context.Context, id uint64, product *models.Product) error
	DeleteProduct(ctx context.Context, id uint64) error
	SearchProducts(ctx context.Context, query string) ([]models.Product, error)
}

type Service struct {
	Product ProductServiceInterface
}

func NewService(repo *repository.Repository, cfg *config.CentrifugConfig) *Service {
	return &Service{
		Product: NewProductService(repo.Product, cfg),
	}
}
