package cart

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain"
)

type Cart struct {
	domain.BaseModel

	CustomerId uint `gorm:"column:customer_id;"`
	OptionId   uint `gorm:"column:option_id;"`
	Quantity   uint `gorm:"column:quantity; type:unsigned big int; not null;"`
}

func (Cart) TableName() string {
	return "cart"
}
