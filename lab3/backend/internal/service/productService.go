package service

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/config"
	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/models"
	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/repository"
)

type productService struct {
	repo repository.ProductRepositoryInterface
	cfg  *config.CentrifugConfig
}

func NewProductService(repo repository.ProductRepositoryInterface, cfg *config.CentrifugConfig) ProductServiceInterface {
	return &productService{
		repo: repo,
		cfg:  cfg,
	}
}

func (s *productService) GetProducts(ctx context.Context) ([]models.Product, error) {
	return s.repo.GetAll(ctx)
}

func (s *productService) CreateProduct(ctx context.Context, product *models.Product) (uint, error) {
	id, err := s.repo.Create(ctx, product)
	if err == nil {
		s.publishToCentrifugo("products", map[string]interface{}{
			"action":  "create",
			"product": product,
		})
	}
	return id, err
}

func (s *productService) publishToCentrifugo(channel string, data interface{}) {
	payload := map[string]interface{}{
		"method": "publish",
		"params": map[string]interface{}{
			"channel": channel,
			"data":    data,
		},
	}
	body, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", s.cfg.CentrifugoURL, bytes.NewBuffer(body))
	if err != nil {
		slog.Error("failed to create centrifugo request", "error", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "apikey "+s.cfg.CentrifugoKey)

	go func() {
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			slog.Error("failed to publish to centrifugo", "error", err)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			slog.Error("centrifugo returned non-ok status", "status", resp.StatusCode)
		}
	}()
}

func (s *productService) GetProductByID(ctx context.Context, id uint64) (*models.Product, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *productService) SearchProducts(ctx context.Context, query string) ([]models.Product, error) {
	return s.repo.Search(ctx, query)
}

func (s *productService) UpdateProduct(ctx context.Context, id uint64, product *models.Product) error {
	return s.repo.Update(ctx, id, product)
}

func (s *productService) DeleteProduct(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}
