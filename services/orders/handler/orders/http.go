package handler

import (
	"net/http"

	util "github.com/RaihanSultana/go-restaurant-management/services/common"
	"github.com/RaihanSultana/go-restaurant-management/services/common/genproto/orders"
	"github.com/RaihanSultana/go-restaurant-management/services/orders/models"
	"github.com/RaihanSultana/go-restaurant-management/services/orders/types"
)

type OrdersHttpHandler struct {
	orderService types.OrderService
	// orders.UnimplementedOrderServiceServer
}

func NewHttpOrderService(orderService types.OrderService) *OrdersHttpHandler {
	handler := &OrdersHttpHandler{
		orderService: orderService,
	}
	return handler

	// // register the OrderServer
	// orders.RegisterOrderServiceServer(grpc, gRPCHandler)
}

func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

// func mapRequestToOrder(res)

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := &models.Order{
		// OrderID:    42,
		CustomerID: 2,
		// ProductID:  1,
		Quantity: 10,
	}

	// order := mapRequestToOrder(req)

	err = h.orderService.CreateOrder(r.Context(), order)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &orders.CreateOrderResponse{Status: "success"}
	util.WriteJSON(w, http.StatusOK, res)
}
