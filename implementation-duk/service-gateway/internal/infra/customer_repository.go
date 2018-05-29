package repository

import (
	"fmt"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/config"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/user"
	"github.com/jinzhu/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(connect *config.DBConnection) user.Repository {
	return &CustomerRepository{connect.GetDB()}
}

func (r *CustomerRepository) FindOne(id int) *user.Customer {
	// TODO : error handling
	customer := &user.Customer{}
	r.db.Where("id = ?", id).First(customer)
	return customer
}

func (r *CustomerRepository) Save(user *user.Customer) {
	err := r.db.Create(user).Error
	if err == nil {
		fmt.Println(err.Error())
	}
	// TODO : error handling
}