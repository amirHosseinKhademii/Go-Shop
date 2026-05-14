package products

import (
	"context"
	repository "shop/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repository.Product, error)
	GetProductById(ctx context.Context, id int32) (repository.Product, error)
}

type svc struct {
	repository repository.Querier
}

// Constructor
func NewService(repository repository.Querier) Service {
	return &svc{repository}
}

func (svc *svc) ListProducts(ctx context.Context) ([]repository.Product, error) {
	return svc.repository.ListProducts(ctx)
}

func (svc *svc) GetProductById(ctx context.Context, id int32) (repository.Product, error) {
	return svc.repository.ProductById(ctx, id)
}
