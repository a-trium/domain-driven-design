package order

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/product"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/user"
)

type Order struct {
	domain.BaseModel

	Customer   *user.Customer `gorm:"foreignkey:CustomerId;"`
	CustomerId uint           `gorm:"column:customer_Id;"`

	//Details []Detail
	//TODO : Status (주문상태)
}

func (Order) TableName() string {
	return "order"
}

type Detail struct {
	domain.BaseModel

	Option   *product.Option `gorm:"foreignkey:OptionId;"`
	OptionId uint            `gorm:"column:option_id;" sql:"type:int REFERENCES Option(id) ON DELETE RESTRICT ON UPDATE CASCADE;"`
	Quantity uint            `gorm:"column:quantity; type:unsigned big int; not null;"`
}

func (Detail) TableName() string {
	return "order_detail"
}
