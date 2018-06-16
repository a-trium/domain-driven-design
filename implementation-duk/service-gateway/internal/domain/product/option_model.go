package product

import "github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain"

type Option struct {
	domain.BaseModel

	ProductId uint   `gorm:"column:product_id; type:unsigned big int;"`
	Name      string `gorm:"column:name; type:varchar(100);"`
	Stock     uint   `gorm:"column:stock; type:unsigned big int; not null;"`
	Price     uint   `gorm:"column:price; type:unsigned big int; not null;"`
}

func (Option) TableName() string {
	return "product_option"
}


