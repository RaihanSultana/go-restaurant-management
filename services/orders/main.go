package main

import (
	"database/sql"
	"log"

	database "github.com/RaihanSultana/go-restaurant-management/services/common/db"
)

var (
	db  *sql.DB
	err error
)

func main() {
	config := database.NewEnvDBConfig()
	db, err = database.ConnectToDB(config)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	defer db.Close()
	log.Println("Database connected successfully")

	httpServer := NewHttpServer(":8000")
	go httpServer.Run()

	grpcServer := NewGRPCServer(":9000")
	go grpcServer.Run()

	select {}
}
