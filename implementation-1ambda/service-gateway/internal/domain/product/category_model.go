package product

import (
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/persistent"
)

type Category struct {
	persistent.BaseModel

	Name        string `gorm:"column:name; type:VARCHAR(255); NOT NULL; UNIQUE;"`
	Description string `gorm:"column:description; type:TEXT; NOT NULL;"`
}

func (Category) TableName() string {
	return "Category"
}
