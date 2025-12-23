package main

import (
	"log"
	"net/http"

	handler "github.com/RaihanSultana/go-restaurant-management/services/orders/handler/orders"
	"github.com/RaihanSultana/go-restaurant-management/services/orders/service"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()
	orderService := service.NewOrderService()
	orderHandler := handler.NewHttpOrderService(orderService)
	orderHandler.RegisterRouter(router)
	log.Println("Starting servcer on: ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
