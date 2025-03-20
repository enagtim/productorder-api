package main

import (
	"order-api/configs"
	"order-api/iternal/product"
	"order-api/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDatabase(conf)
	db.AutoMigrate(&product.Product{})
}
