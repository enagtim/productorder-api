package repository

import (
	"order-api/internal/product/model"
	"order-api/pkg/db"

	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	Database *db.Db
}

func NewProductRepository(db *db.Db) *ProductRepository {
	return &ProductRepository{
		Database: db,
	}
}

func (repo *ProductRepository) Create(product *model.Product) (*model.Product, error) {
	result := repo.Database.DB.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *ProductRepository) GetProducts(limit, offset int) (*[]model.Product, error) {
	var products []model.Product
	result := repo.Database.DB.Limit(limit).Offset(offset).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return &products, nil
}

func (repo *ProductRepository) GetById(id uint) (*model.Product, error) {
	var product model.Product
	result := repo.Database.DB.First(&product, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (repo *ProductRepository) Update(p *model.Product) (*model.Product, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(p)
	if result.Error != nil {
		return nil, result.Error
	}
	return p, nil
}
func (repo *ProductRepository) Delete(id uint) error {
	result := repo.Database.DB.Delete(&model.Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
