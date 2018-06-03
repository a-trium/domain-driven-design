package order

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain"
)

type Order struct {
	domain.BaseModel

	CustomerId uint `gorm:"column:customer_id;"`
	Details    []Detail
	//TODO : Status (주문상태)
}

func (Order) TableName() string {
	return "order"
}

type Detail struct {
	domain.BaseModel

	Order     Order `gorm:"foreignkey:OrderID;"`
	OrderId   uint  `gorm:"column:order_id;`
	ProductId uint  `gorm:"column:product_id;`
	OptionId  uint  `gorm:"column:option_id;`
	Quantity  uint  `gorm:"column:quantity; type:unsigned big int; not null;"`
}

func (Detail) TableName() string {
	return "order_detail"
}
