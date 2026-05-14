package products

import (
	"context"
	repository "shop/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repository.Product, error)
}

type svc struct {
	repository repository.Querier
}

func NewService(repository repository.Querier) Service {
	return &svc{repository}
}

func (s *svc) ListProducts(ctx context.Context) ([]repository.Product, error) {
	return s.repository.ListProducts(ctx)
}
