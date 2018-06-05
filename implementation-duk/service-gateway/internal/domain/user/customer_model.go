package user

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain"
)

type Customer struct {
	domain.BaseModel
	Password
	Address
	Contact
	Name     string `gorm:"column:name; type:varchar(20); not null;"`
	Birthday string `gorm:"column:birthday; type:varchar(10)"`
}

func (Customer) TableName() string {
	return "customer"
}

type Repository interface {
	FindOne(id int) *Customer
	Save(customer *Customer)
}
