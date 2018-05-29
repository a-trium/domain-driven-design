package user

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain"
)

type Seller struct {
	domain.BaseModel
	Password
	Address
	Contact
	Name string `gorm:"column:name; type:varchar(20); not null"`
}

func (Seller) TableName() string {
	return "seller"
}
