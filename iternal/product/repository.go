package product

import (
	"order-api/pkg/db"
)

type ProductRepository struct {
	Database *db.Db
}

func NewProductRepository(db *db.Db) *ProductRepository {
	return &ProductRepository{
		Database: db,
	}
}

func (repo *ProductRepository) Create(product *Product) (*Product, error) {
	result := repo.Database.DB.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}
func (repo *ProductRepository) GetById(id uint) (*Product, error) {
	var product Product
	result := repo.Database.DB.First(&product, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}
