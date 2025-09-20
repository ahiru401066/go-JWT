package db

type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Password string
}

func Create(user *User) error {
	return DB.Create(user).Error
}
