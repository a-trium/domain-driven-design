package user

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Birthday time.Time
	Age      int
	Name     string `gorm:"size:255"` // Default size for string is 255, reset it with this tag
	Num      int    `gorm:"AUTO_INCREMENT"`
}

type Repository interface {
	FindOne(id int) *User
	Save(user *User)
}
