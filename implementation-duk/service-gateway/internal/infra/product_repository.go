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

	record := product.NewProduct()
	err := r.db.First(record, id).Related(&record.Options).Related(&record.Tags, "Tags").Error

	if err != nil {
		return nil, exception.NewProductNotFound()
	}

	return record, nil
}

func (r *ProductRepository) FindByTagName(tagName string) []product.Product {

	record := product.NewTag()
	r.db.Where("name = ?", tagName).First(&record)
	r.db.Model(&record).Related(&record.Products, "Products")

	// primary key 조회가 아니기때문에, Where 조건에 조합으로 들어감.. 즉 아래 쿼리로는 검색할 수 없음
	//record := product.NewTag()
	//r.db.Where("tag.name = ?", tagName).First(&record).Related(&record.Products, "Products")

	return record.Products
}
