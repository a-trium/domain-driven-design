package product

import (
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/persistent"
)

type Image struct {
	persistent.BaseModel

	Name string `gorm:"column:name; type:VARCHAR(255); NOT NULL;"`
	Type string `gorm:"column:type; type:VARCHAR(255); NOT NULL;"`
	Path string `gorm:"column:path; type:TEXT; NOT NULL;"`
}

func (Image) TableName() string {
	return "Image"
}
