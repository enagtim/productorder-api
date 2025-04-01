package main

import (
	"fmt"
	"net/http"
	"order-api/configs"
	"order-api/internal/auth"
	"order-api/internal/order"
	"order-api/internal/product"
	"order-api/internal/user"
	"order-api/migrations"
	"order-api/pkg/db"
	"order-api/pkg/middleware"
)

func main() {
	// App initilization
	conf := configs.LoadConfig()
	db := db.NewDatabase(conf)
	router := http.NewServeMux()
	migrations.Migrate(db)
	middleware.LogInit()

	// Repositoryes
	productRepository := product.NewProductRepository(db)
	userRepository := user.NewUserRepository(db)
	orderRepository := order.NewOrderRepository(db)

	// Services
	productService := product.NewProductService(productRepository)
	authService := auth.NewAuthService(userRepository)
	orderService := order.NewOrderService(orderRepository)

	// Handlers
	product.NewProductHandler(router, productService)
	auth.NewAuthHandler(router, &auth.AuthHandlerDeps{
		AuthService: authService,
		Config:      conf,
	})
	order.NewOrderHandler(router, &order.OrderHandlerDeps{
		OrderService: orderService,
		Config:       conf,
	})
	// Server
	server := http.Server{
		Addr:    ":4000",
		Handler: middleware.LoggingResultRequest(router),
	}

	fmt.Println("Server start on port 4000")
	server.ListenAndServe()
}
