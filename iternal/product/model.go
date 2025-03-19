package product

import "github.com/lib/pq"

type Product struct {
	Id          uint           `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Images      pq.StringArray `json:"images"`
	Price       uint           `json:"price"`
	Discount    uint           `json:"discount"`
}

func NewProduct(name, description string, price, discount uint) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Price:       price,
		Discount:    discount,
	}
}
