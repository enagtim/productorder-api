package product

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string   `gorm:"type:varchar(255);index:name;not null"`
	Description string   `gorm:"type:text; not null"`
	Images      []string `gorm:"type:text[];default:null"`
	Price       float64  `gorm:"type:decimal(10,2);index:price;not null"`
	Discount    float64  `gorm:"type:decimal(10,2);index:discount; default:0; not null"`
}

func NewProduct(name, description string, price, discount float64) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Price:       price,
		Discount:    discount,
	}
}
