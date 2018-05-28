package user

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain"
	"time"
)

type Customer struct {
	domain.BaseModel
	Password
	Address
	Contact
	Name     string `gorm:"type:varchar(20); not null"`
	Birthday time.Time
}

func (Customer) TableName() string {
	return "customer"
}

type Repository interface {
	FindOne(id int) *Customer
	Save(customer *Customer)
}
