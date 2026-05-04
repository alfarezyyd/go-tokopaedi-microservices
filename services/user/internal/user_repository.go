package internal

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(gormTransaction *gorm.DB) ([]User, error)
	Create(gormTransaction *gorm.DB, userEntity *User) error
}
