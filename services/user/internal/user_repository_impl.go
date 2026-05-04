package internal

import (
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (userRepositoryImpl *UserRepositoryImpl) FindAll(gormTransaction *gorm.DB) ([]User, error) {
	var userEntities []User
	err := gormTransaction.Find(&userEntities).Error
	return userEntities, err
}

func (userRepositoryImpl *UserRepositoryImpl) Create(gormTransaction *gorm.DB, userEntity *User) error {
	return gormTransaction.Create(userEntity).Error
}
