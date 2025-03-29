package user

import (
	"order-api/pkg/db"

	"gorm.io/gorm/clause"
)

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{Database: database}
}

func (repo *UserRepository) Create(u *User) (*User, error) {
	result := repo.Database.DB.Create(u)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}
func (repo *UserRepository) Update(u *User) (*User, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(u)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}
func (repo *UserRepository) FindByPhone(phone string) (*User, error) {
	var user User
	result := repo.Database.DB.First(&user, "phone = ?", phone)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
