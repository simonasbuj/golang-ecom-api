package products

import (
	"context"
	"fmt"
	repo "golang-ecom-api/internal/adapters/sqlite/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	GetProductByID(ctx context.Context, id int64) (*repo.Product, error)
}

type service struct{
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) ListProducts(ctx context.Context) ([]repo.Product, error) {
	products, err := s.repo.ListProducts(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get products from database: %w", err)
	}

	return products, nil
}

func (s *service) GetProductByID(ctx context.Context, id int64) (*repo.Product, error) {
	product, err := s.repo.GetProductByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get product from database: %w", err)
	}

	return &product, nil
}
