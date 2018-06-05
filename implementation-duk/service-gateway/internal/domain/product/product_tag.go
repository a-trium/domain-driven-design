package product

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain"
)

type ProductTag struct {
	domain.BaseModel

	ProductId uint `gorm:"product_id; type:unsigned big int; not null; index"`
	TagId     uint `gorm:"tag_id; type:unsigned big int; not null; index"`
}

func (ProductTag) TableName() string {
	return "product_tag"
}
