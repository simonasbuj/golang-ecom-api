package products

import "context"

type Service interface {
	ListProducts(ctx context.Context) error
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) ListProducts(ctx context.Context) error {
	return nil
}
