package product

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model

	Name        string `gorm:"column:name;"`
	Description string `gorm:"column:description;"`
}

func (Category) TableName() string {
	return "Category"
}
