package persistence

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/RaihanSultana/go-restaurant-management/services/orders/models"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, order *models.Order) error {
	fmt.Println("query")
	query := `
		INSERT INTO orders (customer_id, subtotal, total)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.ExecContext(ctx, query, order.CustomerID, order.Subtotal, order.Total)
	fmt.Println("Executed")
	if err != nil {
		fmt.Println("error in executing query", err)
	}

	return err
}
