package service

import (
	"context"
	"fmt"

	"github.com/RaihanSultana/go-restaurant-management/services/common/genproto/orders"
	"github.com/RaihanSultana/go-restaurant-management/services/orders/models"
	"github.com/RaihanSultana/go-restaurant-management/services/orders/persistence"
)

var ordersDb = make([]*orders.Order, 0)

type OrderService struct {
	repo *persistence.OrderRepository
}

func NewOrderService(repo *persistence.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *models.Order) error {
	// ordersDb = append(ordersDb, order)
	fmt.Println("Creating")
	s.repo.CreateOrder(ctx, order)
	return nil
}

func (s *OrderService) GetOrders(ctx context.Context) []*orders.Order {
	return ordersDb
}
