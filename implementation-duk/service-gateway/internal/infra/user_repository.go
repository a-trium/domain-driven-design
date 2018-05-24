package repository

import (
	"fmt"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/config"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/user"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(connect *config.DBConnection) user.Repository {
	return &UserRepository{connect.GetDB()}
}

func (r *UserRepository) FindOne(id int) *user.User {
	// TODO : error handling
	user := &user.User{}
	r.db.Where("id = ?", id).First(user)
	return user
}

func (r *UserRepository) Save(user *user.User) {
	err := r.db.Create(user).Error
	if err == nil {
		fmt.Println(err.Error())
	}
	// TODO : error handling
}