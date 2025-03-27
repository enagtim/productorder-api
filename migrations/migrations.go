package migrations

import (
	"order-api/internal/product"
	"order-api/internal/user"
	"order-api/pkg/db"
)

func Migrate(db *db.Db) {
	err := db.AutoMigrate(&product.Product{}, &user.User{})
	if err != nil {
		panic(err)
	}
}
