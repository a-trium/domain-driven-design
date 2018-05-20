package product

import (
	"github.com/jinzhu/gorm"
)

type Image struct {
	gorm.Model

	Name string `gorm:"column:name; not null;"`
	Type string `gorm:"column:type; not null;"`
	Path string `gorm:"column:path; not null;"`
}

func (Image) TableName() string {
	return "Image"
}
