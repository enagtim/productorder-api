package order

import (
	"order-api/internal/product"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID    uint        `gorm:"index:user_id;not null"`
	Total     float64     `gorm:"type:decimal(10,2);default:0;not null"`
	OrderItem []OrderItem `gorm:"many2many:order_products;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint            `gorm:"index:order_id;not null"`
	Order     Order           `gorm:"constraint:OnDelete:CASCADE;"`
	ProductID uint            `gorm:"index:product_id;not null"`
	Product   product.Product `gorm:"constraint:OnDelete:CASCADE;"`
	Quantity  float64         `gorm:"type:decimal(10,2);default:1;not null"`
}
