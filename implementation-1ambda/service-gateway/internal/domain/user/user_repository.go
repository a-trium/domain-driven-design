package user

import (
	. "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/exception"
)

type UserRepository interface {
	Create(user *User) (Exception)
	Delete(id uint) (bool, Exception)
	FindOne(id uint) (*User, Exception)
	FineAll() (*[]User, Exception)
}
