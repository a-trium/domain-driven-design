package user

import (
	e "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/exception"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Repository interface {
	AddUser(user *User) (*User, e.Exception)
	DeleteUser(id uint) (bool, e.Exception)
	FindUserById(id uint) (*User, e.Exception)
	FineAllUsers() (*[]User, e.Exception)
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) AddUser(record *User) (*User, e.Exception) {
	err := r.db.Create(record).Error

	if err != nil {
		wrap := errors.Wrap(err, "Failed to create User")
		return nil, e.NewInternalServerException(wrap)
	}

	return record, nil
}

func (r *repositoryImpl) DeleteUser(id uint) (bool, e.Exception) {
	record := &User{}
	result := r.db.Where("id = ?", id).Delete(record)

	if result.Error != nil {
		wrap := errors.Wrap(result.Error, "Failed to delete User")
		return false, e.NewInternalServerException(wrap)
	}

	if result.RowsAffected < 1 {
		wrap := errors.Wrap(result.Error, "Failed to delete User")
		return false, e.NewNotFoundException(wrap)
	}

	return true, nil
}

func (r *repositoryImpl) FindUserById(id uint) (*User, e.Exception) {
	record := &User{}
	err := r.db.Where("id = ?", id).First(record).Error

	if err != nil {
		wrap := errors.Wrap(err, "Failed to find-one User")

		if gorm.IsRecordNotFoundError(err) {
			return nil, e.NewNotFoundException(wrap)
		}

		return nil, e.NewInternalServerException(wrap)
	}

	return record, nil
}

func (r *repositoryImpl) FineAllUsers() (*[]User, e.Exception) {
	// TODO: use db.tx

	var records []User

	err := r.db.Find(&records).Error
	if err != nil {
		wrap := errors.Wrap(err, "Failed to find-all User")
		return nil, e.NewInternalServerException(wrap)
	}

	return &records, nil
}
