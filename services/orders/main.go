package main

import (
	"log"

	database "github.com/RaihanSultana/go-restaurant-management/services/common/db"
)

func main() {
	httpServer := NewHttpServer(":8000")
	go httpServer.Run()

	grpcServer := NewGRPCServer(":9000")
	go grpcServer.Run()

	config := database.NewEnvDBConfig()
	db, err := database.ConnectToDB(config)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	defer db.Close()
	log.Println("Database connected successfully")

	select {}
}
