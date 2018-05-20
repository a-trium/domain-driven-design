package product

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model

	Name        string `gorm:"column:name; not null; unique;"`
	Description string `gorm:"column:description; not null;"`
}

func (Category) TableName() string {
	return "Category"
}
