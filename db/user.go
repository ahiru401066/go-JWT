package db

import (
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *User) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &GormUserRepository{DB: db}
}

type User struct {
	ID        uint
	Name      string
	Password  string
	CreatedAt time.Time
}

type GormUserRepository struct {
	DB *gorm.DB
}

func (r *GormUserRepository) Create(user *User) error {
	return DB.Create(user).Error
}
