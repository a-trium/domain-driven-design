package product

import "github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain"

type Tag struct {
	domain.BaseModel
	Name string `gorm:"column:name; type:varchar(20); not null; unique; index;"`

	Products []Product `gorm:"many2many:product_tag;"`
}

func (Tag) TableName() string {
	return "tag"
}

func NewTag() *Tag {
	return &Tag{
		Products: make([]Product, 1),
	}
}
