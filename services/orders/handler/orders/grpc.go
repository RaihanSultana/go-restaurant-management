package handler

import (
	"context"

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

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (resp *orders.CreateOrderResponse, error error) {
	order := &orders.Order{
		OrderID:    42,
		CustomerID: 2,
		ProductID:  1,
		Quantity:   10,
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
