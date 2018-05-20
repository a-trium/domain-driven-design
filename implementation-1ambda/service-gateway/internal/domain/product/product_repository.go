package product

import (
	. "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/exception"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Repository interface {
	AddCategory(record *Category) (*Category, Exception)
	AddImage(record *Image) (*Image, Exception)
	FindCategoryById(id uint) (*Category, Exception)
	FindImageById(id uint) (*Image, Exception)
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}

type repositoryImpl struct {
	db *gorm.DB
}

func (r *repositoryImpl) AddCategory(record *Category) (*Category, Exception) {
	err := r.db.Create(record).Error

	if err != nil {
		wrap := errors.Wrap(err, "Failed to create Category")
		return nil, NewInternalServerException(wrap)
	}

	return record, nil
}

func (r *repositoryImpl) AddImage(record *Image) (*Image, Exception) {
	err := r.db.Create(record).Error

	if err != nil {
		wrap := errors.Wrap(err, "Failed to create Image")
		return nil, NewInternalServerException(wrap)
	}

	return record, nil
}

func (r *repositoryImpl) FindCategoryById(id uint) (*Category, Exception) {
	record := &Category{}
	err := r.db.Where("id = ?", id).First(record).Error

	if err != nil {
		wrap := errors.Wrap(err, "Failed to find Category by id")

		if gorm.IsRecordNotFoundError(err) {
			return nil, NewNotFoundException(wrap)
		}

		return nil, NewInternalServerException(wrap)
	}

	return record, nil
}

func (r *repositoryImpl) FindImageById(id uint) (*Image, Exception) {
	record := &Image{}
	err := r.db.Where("id = ?", id).First(record).Error

	if err != nil {
		wrap := errors.Wrap(err, "Failed to find Image by id")

		if gorm.IsRecordNotFoundError(err) {
			return nil, NewNotFoundException(wrap)
		}

		return nil, NewInternalServerException(wrap)
	}

	return record, nil
}
