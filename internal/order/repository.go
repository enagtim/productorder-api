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

func (repo *OrderRepository) FindOrderByID(orderID, userID uint) (*Order, error) {
	var order Order
	result := repo.Database.DB.Preload("Products").Where("id = ? AND user_id = ?", orderID, userID).First(&order)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

func (repo *OrderRepository) GetAllProductsByUser(userID uint) (*[]Order, error) {
	var orders []Order
	result := repo.Database.DB.Preload("Products").Where("user_id = ?", userID).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return &orders, nil
}
