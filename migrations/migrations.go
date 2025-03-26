package migrations

import (
	"order-api/internal/product/model"
	"order-api/pkg/db"
)

func Migrate(db *db.Db) {
	err := db.AutoMigrate(&model.Product{})
	if err != nil {
		panic(err)
	}
}
