package user

import (
	. "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/exception"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Repository interface {
	CreateUser(user *User) (*User, Exception)
	DeleteUser(id uint) (bool, Exception)
	FindUserById(id uint) (*User, Exception)
	FineAllUsers() (*[]User, Exception)
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) CreateUser(record *User) (*User, Exception) {
	err := r.db.Create(record).Error

	if err != nil {
		wrap := errors.Wrap(err, "Failed to create User")
		return nil, NewInternalServerException(wrap)
	}

	return record, nil
}

func (r *repositoryImpl) DeleteUser(id uint) (bool, Exception) {
	record := &User{}
	result := r.db.Where("id = ?", id).Delete(record)

	if result.Error != nil {
		wrap := errors.Wrap(result.Error, "Failed to delete User")
		return false, NewInternalServerException(wrap)
	}

	if result.RowsAffected < 1 {
		wrap := errors.Wrap(result.Error, "Failed to delete User")
		return false, NewNotFoundException(wrap)
	}

	return true, nil
}

func (r *repositoryImpl) FindUserById(id uint) (*User, Exception) {
	record := &User{}
	err := r.db.Where("id = ?", id).First(record).Error

	if err != nil {
		wrap := errors.Wrap(err, "Failed to find-one User")

		if gorm.IsRecordNotFoundError(err) {
			return nil, NewNotFoundException(wrap)
		}

		return nil, NewInternalServerException(wrap)
	}

	return record, nil
}

func (r *repositoryImpl) FineAllUsers() (*[]User, Exception) {
	// TODO: use db.tx

	var records []User

	err := r.db.Find(&records).Error
	if err != nil {
		wrap := errors.Wrap(err, "Failed to find-all User")
		return nil, NewInternalServerException(wrap)
	}

	return &records, nil
}
