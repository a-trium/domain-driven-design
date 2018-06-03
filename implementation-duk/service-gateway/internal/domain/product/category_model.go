package product

import "github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain"

type Category struct {
	domain.BaseModel

	Name             string   `gorm:"column:name; type:varchar(20); not null; unique"`
	Description      string   `gorm:"column:description; type:varchar(255); not null;"`
	ParentCategory   Category `gorm:"foreignkey:Category;"`
	ParentCategoryId uint     `gorm:"column:parent_id;`
}

func (Category) TableName() string {
	return "category"
}
