package product

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/user"
)

type Product struct {
	domain.BaseModel

	Name  string `gorm:"type:varchar(50); not null;"`
	Price uint   `gorm:"not null;"`

	Category   *Category `gorm:"foreignkey:CategoryID;"`
	CategoryId uint     `gorm:"column:category_id;" sql:"type:UNSIGNED BIG INT REFERENCES Category(id) ON DELETE RESTRICT ON UPDATE CASCADE;"`

	Seller   *user.Seller `gorm:"foreignkey:SellerId;"`
	SellerId uint        `gorm:"column:seller_id;" sql:"type:UNSIGNED BIG INT REFERENCES Seller(id) ON DELETE RESTRICT ON UPDATE CASCADE;"`

	ImageUrl string

	OnSale bool

	//Options []Option
}

func (Product) TableName() string {
	return "product"
}


type Option struct {
	domain.BaseModel

	Product   *Product `gorm:"foreignkey:ProductID;"`
	ProductId uint    `gorm:"column:product_id;" sql:"type:UNSIGNED BIG INT REFERENCES Product(id) ON DELETE RESTRICT ON UPDATE CASCADE"`

	Stock uint

	Name  string
	Price uint
}

func (o *Option) getPrice() uint {
	return o.Price + o.Product.Price
}

func (Option) TableName() string {
	return "product_option"
}
