package product

import (
	"errors"
	"order-api/pkg/db"
)

type ProductRepository struct {
	Db *db.Db
}

func NewProductRepository(db *db.Db) *ProductRepository {
	return &ProductRepository{
		Db: db,
	}
}

func (repo *ProductRepository) Create(p *Product) (*Product, error) {
	query := `INSERT INTO products (name, description, price, discount) 
			  VALUES($1, $2, $3, $4)
			  RETURNING id`
	err := repo.Db.QueryRow(query, p.Name, p.Description, p.Price, p.Discount).Scan(&p.Id)
	if err != nil {
		return nil, errors.New("ERROR SAVE PRODUCT IN PRODUCTS_TABLE")
	}
	return p, nil
}
func (repo *ProductRepository) GetById(id uint) (*Product, error) {
	var p Product
	query := `SELECT id, name, description, images, price, discount 
			  FROM products
			  WHERE id = $1`
	err := repo.Db.QueryRow(query, id).Scan(
		&p.Id,
		&p.Name,
		&p.Description,
		&p.Images,
		&p.Price,
		&p.Discount,
	)
	if err != nil {
		return nil, errors.New("ERROR GET PRODUCT BY ID")
	}
	return &p, nil

}
func (repo *ProductRepository) Update() {

}

func (repo *ProductRepository) Delete() {

}
