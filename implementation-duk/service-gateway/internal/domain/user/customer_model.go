package user

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Customer struct {
	gorm.Model
	Birthday time.Time
	Name     string `gorm:"size:20"` // Default size for string is 255, reset it with this tag
}

type Repository interface {
	FindOne(id int) *Customer
	Save(customer *Customer)
}
