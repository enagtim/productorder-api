package order

import (
	"order-api/internal/product"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID   uint              `gorm:"index:user_id;not null"`
	Products []product.Product `gorm:"many2many:order_products;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
