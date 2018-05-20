package product

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model

	Name   string `gorm:"column:name; not null;"`
	Price  string `gorm:"column:price; not null;"`
	Detail string `gorm:"column:detail; not null;"`

	Category    Category `gorm:"foreignkey:CategoryRef"`
	CategoryRef uint

	Image    Image `gorm:"foreignkey:ImageRef"`
	ImageRef uint
}

func (Product) TableName() string {
	return "Product"
}
