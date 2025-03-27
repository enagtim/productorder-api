package user

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Phone     string `gorm:"type:varchar(11);uniqueIndex;not null"`
	SessionId string `gorm:"type:varchar(50);not null"`
}
