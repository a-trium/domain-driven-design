package repository

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/config/db"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/product"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/exception"
	"github.com/jinzhu/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(datasource db.DataSource) product.Repository {
	return &ProductRepository{datasource.GetConnection()}
}

func (r *ProductRepository) FindById(id int) (*product.Product, error) {

	record := product.New()
	err := r.db.First(&record, id).Error

	if err != nil {
		return nil, exception.NewProductNotFound()
	}

	r.db.Model(record).Related(&record.Options)
	r.db.Model(record).Related(&record.Tags)

	return record, nil
}

func (r *ProductRepository) FindByTagId(tagId int) []product.Product {

	//products := make([]product.Product, 1)
	//r.db.Where("name = ?", tagId)

	return nil
}
