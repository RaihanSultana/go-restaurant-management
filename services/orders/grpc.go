package main

import (
	"log"
	"net"

	handler "github.com/RaihanSultana/go-restaurant-management/services/orders/handler/orders"
	"github.com/RaihanSultana/go-restaurant-management/services/orders/persistence"
	"github.com/RaihanSultana/go-restaurant-management/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	//register your grpc service
	repo := persistence.NewOrderRepository(db)
	orderService := service.NewOrderService(repo)
	handler.NewGrpcOrderService(grpcServer, orderService)

	log.Println("starting grpc server on:", s.addr)
	return grpcServer.Serve(lis)
}
