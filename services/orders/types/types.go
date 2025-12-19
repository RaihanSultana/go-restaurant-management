package types

import (
	"context"

	"github.com/RaihanSultana/go-restaurant-management/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
}
