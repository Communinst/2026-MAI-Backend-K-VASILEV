package service

import (
	"context"

	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/models"
	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/repository"
)

type productService struct {
	repo repository.ProductRepositoryInterface
}

func NewProductService(repo repository.ProductRepositoryInterface) ProductServiceInterface {
	return &productService{
		repo: repo,
	}
}

func (s *productService) GetProducts(ctx context.Context) ([]models.Product, error) {
	return s.repo.GetAll(ctx)
}

func (s *productService) CreateProduct(ctx context.Context, product *models.Product) (uint, error) {
	return s.repo.Create(ctx, product)
}

func (s *productService) GetProductByID(ctx context.Context, id uint64) (*models.Product, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *productService) SearchProducts(ctx context.Context, query string) ([]models.Product, error) {
	return s.repo.Search(ctx, query)
}
