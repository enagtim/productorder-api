package migrations

import (
	"order-api/internal/order"
	"order-api/internal/product"
	"order-api/internal/user"
	"order-api/pkg/db"
)

func Migrate(db *db.Db) {
	err := db.AutoMigrate(&product.Product{}, &user.User{}, &order.Order{}, &order.OrderItem{})
	if err != nil {
		panic(err)
	}
}
