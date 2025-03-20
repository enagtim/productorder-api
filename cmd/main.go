package main

import (
	"fmt"
	"net/http"
	"order-api/configs"
	"order-api/iternal/product"
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
	productRepository := product.NewProductRepository(db)

	// Handlers
	product.NewProductHandler(router, &product.ProductHandler{
		ProductRepository: productRepository,
	})

	// Server
	server := http.Server{
		Addr:    ":4000",
		Handler: router,
	}
	fmt.Println("Server start on port 4000")
	server.ListenAndServe()
}
