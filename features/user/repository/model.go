package repository

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(15)"`
	Email    string
	Phone    string
	Address  string
	Role     string
}
