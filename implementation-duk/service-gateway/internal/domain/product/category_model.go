package product

import "github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain"

type Category struct {
	domain.BaseModel

	Name        string `gorm:"type:varchar(20); not null; unique"`
	Description string

	//SubCategory   Category
	//SubCategoryId Category
}
