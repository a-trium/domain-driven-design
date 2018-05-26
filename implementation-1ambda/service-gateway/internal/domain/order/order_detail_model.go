package order

import (
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/persistent"
)

type OrderDetail struct {
	persistent.BaseModel

	Index    uint `gorm:"column:index; 		type:UNSIGNED BIG INT; 	NOT NULL;"`
	Price    uint `gorm:"column:price; 		type:UNSIGNED BIG INT; 	NOT NULL;"`
	Quantity uint `gorm:"column:quantity; 	type:UNSIGNED BIG INT; 	NOT NULL;"`
	Amount   uint `gorm:"column:amount; 	type:UNSIGNED BIG INT; 	NOT NULL;"`

	ProductID uint            `gorm:"column:product_id;" sql:"type:UNSIGNED BIG INT REFERENCES Product(id) ON DELETE RESTRICT ON UPDATE CASCADE;"`
}

func (OrderDetail) TableName() string {
	return "OrderDetail"
}
