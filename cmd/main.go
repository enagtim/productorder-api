package main

import (
	"fmt"
	"net/http"
	"order-api/configs"
	"order-api/internal/product"
	"order-api/internal/product/repository"
	"order-api/internal/product/service"
	"order-api/migrations"
	"order-api/pkg/db"
)

func main() {

	// Configs, Database, Router
	conf := configs.LoadConfig()
	db := db.NewDatabase(conf)
	router := http.NewServeMux()

	// Database Migrations
	migrations.Migrate(db)

	// Repositoryes
	productRepository := repository.NewProductRepository(db)

	// Services

	service := service.NewProductService(productRepository)

	// Handlers
	product.NewProductHandler(router, service)

	// Server
	server := http.Server{
		Addr:    ":4000",
		Handler: router,
	}
	fmt.Println("Server start on port 4000")
	server.ListenAndServe()
}
