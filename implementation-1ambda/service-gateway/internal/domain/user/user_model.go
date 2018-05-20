package user

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Password string `gorm:"column:password; not null;"`
	Email    string `gorm:"column:email; not null;"`
	Phone    string `gorm:"column:phone; not null;"`
	Name     string `gorm:"column:name; not null;"`
	Birthday string `gorm:"column:birthday; not null;'"`
	Address  string `gorm:"column:address; not null;"`
}

func (User) TableName() string {
	return "User"
}



