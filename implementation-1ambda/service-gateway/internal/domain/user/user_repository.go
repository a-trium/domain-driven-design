package user

import (
	. "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/exception"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type UserRepository interface {
	Create(user *User) (Exception)
	Delete(id uint) (bool, Exception)
	FindOne(id uint) (*User, Exception)
	FineAll() (*[]User, Exception)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) Create(record *User) Exception {
	err := r.db.Create(record).Error

	if err != nil {
		wrap := errors.Wrap(err, "Failed to create User")
		return NewInternalServerException(wrap)
	}

	return nil
}

func (r *userRepositoryImpl) Delete(id uint) (bool, Exception) {
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

func (r *userRepositoryImpl) FindOne(id uint) (*User, Exception) {
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

func (r *userRepositoryImpl) FineAll() (*[]User, Exception) {
	// TODO: use db.tx

	var records []User

	err := r.db.Find(&records).Error
	if err != nil {
		wrap := errors.Wrap(err, "Failed to find-all User")
		return nil, NewInternalServerException(wrap)
	}

	return &records, nil
}
