package order

import (
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/domain/product"
	"github.com/jinzhu/gorm"
)

type OrderDetail struct {
	gorm.Model

	Index    uint `gorm:"column:index; not null;"`
	Price    uint `gorm:"column:price; not null;"`
	Quantity uint `gorm:"column:quantity; not null;"`
	Amount   uint `gorm:"column:amount; not null;"`

	Order      Order           `gorm:"foreignkey:OrderRef"`
	OrderRef   uint
	Product    product.Product `gorm:"foreignkey:ProductRef"`
	ProductRef uint
}

func (OrderDetail) TableName() string {
	return "OrderDetail"
}
