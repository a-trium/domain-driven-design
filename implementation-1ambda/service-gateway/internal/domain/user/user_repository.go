package user

import (
	"strings"

	e "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/exception"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type AuthCookie string

type Repository interface {
	AddUser(user *User) (*User, e.Exception)
	DeleteUser(id uint) (bool, e.Exception)
	FindUserById(id uint) (*User, e.Exception)
	FineAllUsers() (*[]User, e.Exception)

	Register(uid string, password string) (*AuthIdentity, e.Exception)
	Login(uid string, password string) (AuthCookie, e.Exception)
	Logout(uid string) (e.Exception)
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

func (r *repositoryImpl) Register(uid string, password string) (*AuthIdentity, e.Exception) {
	if strings.TrimSpace(uid) == "" || strings.TrimSpace(password) == "" {
		err := errors.New("Empty uid or password")
		return nil, e.NewBadRequestException(err)
	}

	tx := r.db.Begin()

	user := User{}
	err := tx.Create(&user).Error

	if err != nil {
		tx.Rollback()
		wrap := errors.Wrap(err, "Failed to create User")
		return nil, e.NewInternalServerException(wrap)
	}

	authIdentity := &AuthIdentity{
		Provider: ProviderPassword,
		UID: uid,
		EncryptedPassword: password,

		User: user,
	}

	err = tx.Create(authIdentity).Error
	if err != nil {
		tx.Rollback()
		wrap := errors.Wrap(err, "Failed to create AuthIdentity")
		return nil, e.NewInternalServerException(wrap)
	}

	tx.Commit()
	return authIdentity, nil
}

func (r *repositoryImpl) Login(uid string, password string) (AuthCookie, e.Exception) {
	return AuthCookie(""), nil
}

func (r *repositoryImpl) Logout(uid string) (e.Exception) {
	return nil
}
