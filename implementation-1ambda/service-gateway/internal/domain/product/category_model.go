package product

import (
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/persistent"
)

type Category struct {
	persistent.BaseModel

	Path        string `gorm:"column:path; type:VARCHAR(255); NOT NULL; UNIQUE;"`
	Name        string `gorm:"column:name; type:VARCHAR(255); NOT NULL; INDEX;"`
	DisplayName string `gorm:"column:display_name; type:VARCHAR(255); NOT NULL;"`
	Description string `gorm:"column:description; type:TEXT; NOT NULL;"`

	ParentCategoryID uint `gorm:"column:parent_category_id;" sql:"type:UNSIGNED BIG INT REFERENCES Category(id) ON DELETE NULL ON UPDATE CASCADE;"`
}

func (Category) TableName() string {
	return "Category"
}
