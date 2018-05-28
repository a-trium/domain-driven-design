package cart

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/product"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/user"
)

type Cart struct {
	domain.BaseModel

	Customer user.Customer `gorm:"foreignkey:CustomerId;"`
	CustomerId uint `gorm:"column:customer_Id;"`

	Option product.Option `gorm:"foreignkey:OptionId;"`
	OptionId uint `gorm:"column:option_id;" sql:"type:UNSIGNED BIG INT REFERENCES Option(id) ON DELETE RESTRICT ON UPDATE CASCADE;"`

	Quantity uint
}
