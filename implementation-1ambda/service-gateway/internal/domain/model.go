package domain

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	// TODO encrypt
	Password string `gorm:"column:password; type:varchar(191); not null;"`
	Email    string `gorm:"column:email; type:varchar(191); unique; not null;"`
	Name     string `gorm:"column:name; type:varchar(191); not null;"`
	Birthday string `gorm:"column:birthday; type:varchar(191); not null;"`
	Address  string `gorm:"column:address; type:varchar(191); not null;"`

	//MemberNumber uint `gorm:"column:member_number;AUTO_INCREMENT"`
	// avatar bin image
	// member grade string
}

func (User) TableName() string {
	return "User"
}

type Category struct {
	gorm.Model

	Name string `gorm:"column:name; type:varchar(191); unique; not null;"`
}


func (Category) TableName() string {
	return "Category"
}

type Product struct {
	gorm.Model


}

func (Product) TableName() string {
	return "Product"
}


