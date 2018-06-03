package product

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain"
)

type Product struct {
	domain.BaseModel

	Name     string `gorm:"type:varchar(50); not null;"`
	Price    uint   `gorm:"column:price; type:unsigned big int; not null;"`
	SellerId uint   `gorm:"column:seller_id; type:unsigned big int;"`
	ImageUrl string `gorm:"column:image_url; type:varchar(255);"`
	OnSale   bool   `gorm:"column:on_sale;"`

	Options []Option
	Tags    []Tag
}

func (Product) TableName() string {
	return "product"
}

type Option struct {
	domain.BaseModel

	Product   *Product `gorm:"foreignkey:ProductID;"`
	ProductId uint     `gorm:"column:product_id; type:unsigned big int;"`
	Name      string   `gorm:"column:name; type:varchar(100);"`
	Stock     uint     `gorm:"column:stock; type:unsigned big int; not null;"`
	Price     uint     `gorm:"column:price; type:unsigned big int; not null;"`
}

func (o *Option) getPrice() uint {
	return o.Price + o.Product.Price
}

func (Option) TableName() string {
	return "product_option"
}
