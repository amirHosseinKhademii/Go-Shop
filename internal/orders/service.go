package orders

import (
	"context"
	"fmt"

	repository "shop/internal/adapters/postgresql/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrProductNotFound      = fmt.Errorf("product not found")
	ErrInsufficientQuantity = fmt.Errorf("insufficient quantity")
)

type Service interface {
	GetOrders(ctx context.Context) ([]repository.ListOrdersRow, error)
	CreateOrder(ctx context.Context, req CreateOrderRequest) error
}

type svc struct {
	repo repository.Querier
	pool *pgxpool.Pool
}

func NewService(repo repository.Querier, pool *pgxpool.Pool) Service {
	return &svc{repo, pool}
}

func (s *svc) GetOrders(ctx context.Context) ([]repository.ListOrdersRow, error) {
	return s.repo.ListOrders(ctx)
}

func (s *svc) CreateOrder(ctx context.Context, req CreateOrderRequest) error {
	if req.CustomerID == 0 {
		return fmt.Errorf("customerId is required and must be positive")
	}
	if len(req.Items) == 0 {
		return fmt.Errorf("items is required and must contain at least one item")
	}
	for i, item := range req.Items {
		if item.ProductID == 0 {
			return fmt.Errorf("items[%d].productId must be positive", i)
		}
		if item.Quantity == 0 {
			return fmt.Errorf("items[%d].quantity must be positive", i)
		}
	}

	// create order transaction
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	qtx := repository.New(tx)

	order, err := qtx.CreateOrder(ctx, req.CustomerID)
	if err != nil {
		return err
	}

	for _, item := range req.Items {
		// check if product exists
		product, err := qtx.ProductById(ctx, item.ProductID)
		if err != nil {
			return ErrProductNotFound
		}

		if product.Quantity < item.Quantity {
			return ErrInsufficientQuantity
		}

		_, err = qtx.AddOrderItem(ctx, repository.AddOrderItemParams{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		})
		if err != nil {
			return err
		}

		// update product quantity
		newQuantity := product.Quantity - item.Quantity
		err = qtx.UpdateProductQuantity(ctx, repository.UpdateProductQuantityParams{
			ID:       item.ProductID,
			Quantity: newQuantity,
		})
		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}
