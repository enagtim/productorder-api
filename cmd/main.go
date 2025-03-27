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
	"order-api/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDatabase(conf)
	router := http.NewServeMux()

	migrations.Migrate(db)

	productRepository := repository.NewProductRepository(db)
	service := service.NewProductService(productRepository)

	product.NewProductHandler(router, service)

	middleware.LogInit()

	server := http.Server{
		Addr:    ":4000",
		Handler: middleware.LoggingResultRequest(router),
	}

	fmt.Println("Server start on port 4000")
	server.ListenAndServe()
}
