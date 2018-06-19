package product

import (
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/persistent"
	dto "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/pkg/generated/swagger/swagmodel"
	"strconv"
)

type OnSale string

const OnSaleY OnSale = "Y"
const OnSaleN OnSale = "N"

type Product struct {
	persistent.BaseModel

	Name        string `gorm:"column:name; type:VARCHAR(255); NOT NULL;"`
	Price       uint   `gorm:"column:price; type:UNSIGNED BIG INT; NOT NULL;"`
	Description string `gorm:"column:description; type:TEXT; NOT NULL;"`
	OnSale    OnSale  `gorm:"column:on_sale; type:VARCHAR(4); NOT NULL;"`

	Category   Category `gorm:"foreignkey:CategoryID;"`
	CategoryID uint     `gorm:"column:category_id;" sql:"type:UNSIGNED BIG INT REFERENCES Category(id) ON DELETE RESTRICT ON UPDATE CASCADE;"`

	Image   Image `gorm:"foreignkey:ImageID;"`
	ImageID uint  `gorm:"column:image_id;" sql:"type:UNSIGNED BIG INT NULL REFERENCES Image(id) ON DELETE SET NULL ON UPDATE CASCADE"`
}

func (Product) TableName() string {
	return "Product"
}

func (p *Product) convertToDTO() *dto.Product {
	return &dto.Product{
		CategoryDisplayName: p.Category.DisplayName,
		CategoryID: strconv.FormatUint(uint64(p.CategoryID), 10),
		CategoryPath: p.Category.Path,
		Description: p.Description,
		ID: strconv.FormatUint(uint64(p.ID), 10),
		ImageID: strconv.FormatUint(uint64(p.ImageID), 10),
		ImagePath: p.Image.Path,
		ImageType: p.Image.Type,
		Name: p.Name,
		OnSale: string(p.OnSale),
		Price: strconv.FormatUint(uint64(p.Price), 10),
	}
}
