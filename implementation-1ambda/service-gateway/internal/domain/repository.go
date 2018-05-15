package domain

import (
	. "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/exception"
)

type UserRepository interface {
	create(user *User) (*User, *Error)
	delete(id uint) (bool, *Error)
	findOne(id uint) (*User, *Error)
	fineAll() (*[]User, *Error)
}
