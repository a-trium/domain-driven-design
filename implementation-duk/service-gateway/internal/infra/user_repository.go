package repository

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/user"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	repository user.Repository
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepository{db: db}
}

func (r *userRepository) FindOne(id int) *user.User {

	// TODO : error handling
	user := &user.User{}
	r.db.Where("id = ?", id).First(user)
	return user
}