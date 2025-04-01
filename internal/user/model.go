package user

import (
	"order-api/internal/order"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Phone     string      `gorm:"type:varchar(11);uniqueIndex;not null"`
	SessionId string      `gorm:"type:varchar(50);not null"`
	Code      uint        `gorm:"type:smallint;not null"`
	Order     order.Order `gorm:"foreignKey:UserID"`
}
