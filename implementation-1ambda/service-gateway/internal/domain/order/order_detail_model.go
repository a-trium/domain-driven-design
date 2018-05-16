package order

import (
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/domain/product"
	"github.com/jinzhu/gorm"
)

type OrderDetail struct {
	gorm.Model

	Index    uint `gorm:"column:index;"`
	Price    uint `gorm:"column:price;"`
	Quantity uint `gorm:"column:quantity;"`
	Amount   uint `gorm:"column:amount;"`

	Order      Order           `gorm:"foreignkey:OrderRef"`
	OrderRef   uint
	Product    product.Product `gorm:"foreignkey:ProductRef"`
	ProductRef uint
}

func (OrderDetail) TableName() string {
	return "OrderDetail"
}
