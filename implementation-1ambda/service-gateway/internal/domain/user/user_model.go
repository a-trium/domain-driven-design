package user

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Password string `gorm:"column:password;"`
	Email    string `gorm:"column:email;"`
	Phone    string `gorm:"column:phone;"`
	Name     string `gorm:"column:name;"`
	Birthday string `gorm:"column:birthday;"`
	Address  string `gorm:"column:address;"`
}

func (User) TableName() string {
	return "User"
}



