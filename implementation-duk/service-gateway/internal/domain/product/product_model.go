package product

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain"
	"strings"
)

type Product struct {
	domain.BaseModel

	Name     string `gorm:"type:varchar(50); not null;"`
	Price    uint   `gorm:"column:price; type:unsigned big int; not null;"`
	SellerId uint   `gorm:"column:seller_id; type:unsigned big int;"`
	ImageUrl string `gorm:"column:image_url; type:varchar(255);"`
	Sale     string `gorm:"column:on_sale; type:varchar(2);"`

	Options []Option `gorm:"foreignkey:ProductId"`
	Tags    []Tag    `gorm:"many2many:product_tag;"`
}

func NewProduct() *Product {
	return &Product{
		Options: make([]Option, 1),
		Tags:    make([]Tag, 1),
	}
}

func (p Product) OnSale() bool {
	return strings.EqualFold(p.Sale, "Y")
}

func (Product) TableName() string {
	return "product"
}
