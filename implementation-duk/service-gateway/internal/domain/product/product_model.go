package product

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/user"
)

type Product struct {
	domain.BaseModel

	Name  string `gorm:"type:varchar(50); not null;"`
	Price uint   `gorm:"column:price; type:unsigned big int; not null;"`

	Category   *Category    `gorm:"foreignkey:CategoryID;"`
	CategoryId uint         `gorm:"column:category_id;" sql:"type:int REFERENCES Category(id) ON DELETE RESTRICT ON UPDATE CASCADE;"`
	Seller     *user.Seller `gorm:"foreignkey:SellerId;"`
	SellerId   uint         `gorm:"column:seller_id;" sql:"type:int REFERENCES Seller(id) ON DELETE RESTRICT ON UPDATE CASCADE;"`
	ImageUrl   string       `gorm:"column:image_url;type:varchar(255); "`
	OnSale     bool         `gorm:"column:on_sale;"`

	//Options []Option
}

func (Product) TableName() string {
	return "product"
}

type Option struct {
	domain.BaseModel

	Product   *Product `gorm:"foreignkey:ProductID;"`
	ProductId uint     `gorm:"column:product_id;" sql:"type:int REFERENCES Product(id) ON DELETE RESTRICT ON UPDATE CASCADE"`
	Name      string   `gorm:"column:name; type:varchar(100)"`
	Stock     uint     `gorm:"column:stock; type:unsigned big int; not null;"`
	Price     uint     `gorm:"column:price; type:unsigned big int; not null;"`
}

func (o *Option) getPrice() uint {
	return o.Price + o.Product.Price
}

func (Option) TableName() string {
	return "product_option"
}
