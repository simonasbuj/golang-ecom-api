package orders

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	repo "golang-ecom-api/internal/adapters/sqlite/sqlc"
	"log/slog"
)

var (
	errProductNotFind = errors.New("product not found")
	errProductNoStock = errors.New("remaining item quantity is less than ordered")
)

type Service interface {
	PlaceOrder(ctx context.Context, orderReq createOrderParams) (*repo.Order, error)
}

type service struct {
	repo *repo.Queries
	db   *sql.DB
}

func NewService(repo *repo.Queries, db *sql.DB) Service {
	return &service{
		repo: repo,
		db:   db,
	}
}

func (s *service) PlaceOrder(ctx context.Context, orderReq createOrderParams) (*repo.Order, error) {
	slog.Info("placing order", "order", orderReq)

	if orderReq.CustomerID == 0 {
		return nil, fmt.Errorf("CustomerID is required")
	}

	if len(orderReq.Items) == 0 {
		return nil, fmt.Errorf("at leat one item in order is required")
	}

	tx, err := s.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin db transaction: %w", err)
	}
	defer tx.Rollback()

	qtx := s.repo.WithTx(tx)

	createdOrder, err := qtx.CreateOrder(ctx, orderReq.CustomerID)
	if err != nil {
		return nil, fmt.Errorf("failed to insert order into db: %w", err)
	}

	for _, item := range orderReq.Items {
		product, err := qtx.GetProductByID(ctx, item.ProductID)
		if err != nil {
			return nil, errProductNotFind
		}

		if product.Quantity < item.Quantity {
			return nil, errProductNoStock
		}

		_, err = qtx.CreateOrderItem(ctx, repo.CreateOrderItemParams{
			OrderID:      createdOrder.ID,
			ProductID:    item.ProductID,
			Quantity:     item.Quantity,
			PriceInCents: product.PriceInCents,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to insert order item: %w", err)
		}

		err = qtx.UpdateProductQuantity(ctx, repo.UpdateProductQuantityParams{
			Quantity: product.Quantity - item.Quantity,
			ID:       item.ProductID,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to update product quantity: %w", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit db transaction: %w", err)
	}

	return &createdOrder, nil
}
