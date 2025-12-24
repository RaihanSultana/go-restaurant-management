package handler

import (
	"context"
	"math/rand/v2"

	"github.com/RaihanSultana/go-restaurant-management/services/common/genproto/orders"
	"github.com/RaihanSultana/go-restaurant-management/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	//unimplemented UnimplementedOrderServiceHandler
	orderService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrderService(grpc *grpc.Server, orderService types.OrderService) {
	gRPCHandler := &OrdersGrpcHandler{
		orderService: orderService,
	}

	// register the OrderServer
	orders.RegisterOrderServiceServer(grpc, gRPCHandler)
}

func (h *OrdersGrpcHandler) GetOrders(ctx context.Context, req *orders.GetOrderRequest) (*orders.GetOrderResponse, error) {
	o := h.orderService.GetOrders(ctx)
	res := &orders.GetOrderResponse{
		Orders: o,
	}
	return res, nil
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (resp *orders.CreateOrderResponse, error error) {
	order := &orders.Order{
		OrderID:    rand.Int32(),
		CustomerID: req.CustomerID,
		ProductID:  req.ProductID,
		Quantity:   req.Quantity,
	}

	err := h.orderService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	resp = &orders.CreateOrderResponse{
		Status: "success",
	}

	return resp, nil
}
