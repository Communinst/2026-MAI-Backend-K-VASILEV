package postgres

import (
	"context"

	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/models"
	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/repository"
	"gorm.io/gorm"
)

type ProductPostgresRepository struct {
	db *gorm.DB
}

func NewProductPostgresRepository(db *gorm.DB) repository.ProductRepositoryInterface {
	return &ProductPostgresRepository{
		db: db,
	}
}

func (r *ProductPostgresRepository) GetAll(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	err := r.db.WithContext(ctx).Preload("Category").Find(&products).Error
	return products, err
}

func (r *ProductPostgresRepository) Create(ctx context.Context, product *models.Product) (uint, error) {
	err := r.db.WithContext(ctx).Create(product).Error
	return product.ID, err
}

func (r *ProductPostgresRepository) GetByID(ctx context.Context, id uint64) (*models.Product, error) {
	var product models.Product
	err := r.db.WithContext(ctx).Preload("Category").First(&product, id).Error
	return &product, err
}

func (r *ProductPostgresRepository) Search(ctx context.Context, query string) ([]models.Product, error) {
	var products []models.Product
	searchQuery := "%" + query + "%"
	err := r.db.WithContext(ctx).
		Preload("Category").
		Where("name ILIKE ? OR description ILIKE ?", searchQuery, searchQuery).
		Find(&products).Error
	return products, err
}
