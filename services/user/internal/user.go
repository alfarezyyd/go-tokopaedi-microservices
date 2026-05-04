package internal

import "time"

type User struct {
	Id          uint64    `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"column:name"`
	Email       string    `gorm:"column:email"`
	Password    string    `gorm:"column:password"`
	PhoneNumber string    `gorm:"column:phone_number"`
	ProfilePath string    `gorm:"column:profile_path"`
	IsActive    bool      `gorm:"column:is_active"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}
