package user


import (
	"github.com/jinzhu/gorm"
)

const USER_TABLE = "user"

type User struct {
	gorm.Model

	// TODO encrypt
	Password string `gorm:"column:password;not null;"`
	Email    string `gorm:"column:email;not null;"`
	Name     string `gorm:"column:name;not null;"`
	Birthday string `gorm:"column:birthday;not null;"`
	Address  string `gorm:"column:address;not null;"`

	//MemberNumber uint `gorm:"column:member_number;AUTO_INCREMENT"`
	// avatar bin image
	// member grade string
}
