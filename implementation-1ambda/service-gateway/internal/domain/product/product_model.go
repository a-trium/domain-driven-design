package product

import (
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/persistent"
)

type Product struct {
	persistent.BaseModel

	Name        string `gorm:"column:name; type:VARCHAR(255); NOT NULL;"`
	Price       uint   `gorm:"column:price; type:UNSIGNED BIG INT; NOT NULL;"`
	Description string `gorm:"column:description; type:TEXT; NOT NULL;"`

	Category   Category `gorm:"foreignkey:CategoryID;"`
	CategoryID uint     `gorm:"column:category_id;" sql:"type:UNSIGNED BIG INT REFERENCES Category(id) ON DELETE RESTRICT ON UPDATE CASCADE;"`

	Image   Image `gorm:"foreignkey:ImageID;"`
	ImageID uint  `gorm:"column:image_id;" sql:"type:UNSIGNED BIG INT NULL REFERENCES Image(id) ON DELETE SET NULL ON UPDATE CASCADE"`
}

func (Product) TableName() string {
	return "Product"
}
