package migrations

import (
	"order-api/internal/product"
	"order-api/pkg/db"
)

func Migrate(db *db.Db) {
	err := db.AutoMigrate(&product.Product{})
	if err != nil {
		panic(err)
	}
}
