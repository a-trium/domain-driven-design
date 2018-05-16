package product

import (
	"github.com/jinzhu/gorm"
)

type Image struct {
	gorm.Model

	Name string `gorm:"column:name;"`
	Type string `gorm:"column:type;"`
	Path string `gorm:"column:path;"`
}

func (Image) TableName() string {
	return "Image"
}
