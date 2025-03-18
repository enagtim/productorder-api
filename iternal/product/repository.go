package product

import "order-api/pkg/db"

type ProductRepository struct {
	Database *db.Db
}

func NewProductRepository(db *db.Db) *ProductRepository {
	return &ProductRepository{
		Database: db,
	}
}

func (repo *ProductRepository) Create() {

}
func (repo *ProductRepository) Get() {

}
func (repo *ProductRepository) Update() {

}

func (repo *ProductRepository) Delete() {

}
