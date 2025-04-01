package order

import (
	"errors"
	"order-api/internal/product"
	"order-api/pkg/db"
)

type OrderRepository struct {
	Database *db.Db
}

func NewOrderRepository(db *db.Db) *OrderRepository {
	return &OrderRepository{Database: db}
}

func (repo *OrderRepository) CreateOrder(userID uint, productIDs []uint) (*Order, error) {
	var products []product.Product
	result := repo.Database.DB.Where("id IN ?", productIDs).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(products) == 0 {
		return nil, errors.New("no products found")
	}
	order := Order{
		UserID:   userID,
		Products: products,
	}
	result = repo.Database.DB.Create(&order)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}
