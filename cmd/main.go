package main

import (
	"fmt"
	"net/http"
	"order-api/configs"
	"order-api/iternal/product"
	"order-api/pkg/db"
)

func main() {

	// Configs
	conf := configs.LoadConfig()
	db := db.NewDatabase(conf)
	router := http.NewServeMux()

	// Repository
	productRepository := product.NewProductRepository(db)

	// Handler
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
