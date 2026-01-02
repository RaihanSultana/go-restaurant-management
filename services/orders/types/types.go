package types

import (
	"context"

	"github.com/RaihanSultana/go-restaurant-management/services/common/genproto/orders"
	"github.com/RaihanSultana/go-restaurant-management/services/orders/models"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order *models.Order) error
	GetOrders(context.Context) []*orders.Order
}
