package user

import (
	"strings"

	e "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/exception"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Repository interface {
	DeleteUser(id uint) (bool, e.Exception)
	FindUserById(id uint) (*User, e.Exception)
	FineAllUsers() (*[]User, e.Exception)

	Register(uid string, password string) (*AuthIdentity, e.Exception)
	Authenticate(uid string, password string) (*AuthClaim, e.Exception)
}

type repositoryImpl struct {
	db        *gorm.DB
	encryptor Encryptor
}

func NewRepository(db *gorm.DB, encryptor Encryptor) Repository {
	return &repositoryImpl{db: db, encryptor: encryptor}
}

func (r *repositoryImpl) DeleteUser(id uint) (bool, e.Exception) {
	record := &User{}
	result := r.db.Where("id = ?", id).Delete(record)

	if result.Error != nil {
		wrap := errors.Wrap(result.Error, "Failed to delete User")
		return false, e.NewInternalServerException(wrap)
	}

	if result.RowsAffected < 1 {
		wrap := errors.Wrap(result.Error, "Failed to fine User to be deleted")
		return false, e.NewNotFoundException(wrap)
	}

	return true, nil
}

func (r *repositoryImpl) FindUserById(id uint) (*User, e.Exception) {
	record := &User{}
	err := r.db.Where("id = ?", id).First(record).Error

	if err != nil {
		wrap := errors.Wrap(err, "Failed to find User")

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
		wrap := errors.Wrap(err, "Failed to find all User")
		return nil, e.NewInternalServerException(wrap)
	}

	return &records, nil
}

func (r *repositoryImpl) Register(uid string, password string) (*AuthIdentity, e.Exception) {
	if strings.TrimSpace(uid) == "" || strings.TrimSpace(password) == "" {
		err := errors.New("Empty uid or password")
		return nil, e.NewBadRequestException(err)
	}

	digested, err := r.encryptor.Digest(password)
	if err != nil {
		wrap := errors.Wrap(err, "Failed to digest password")
		return nil, e.NewInternalServerException(wrap)
	}

	tx := r.db.Begin()

	user := User{}
	err = tx.Create(&user).Error

	if err != nil {
		tx.Rollback()
		wrap := errors.Wrap(err, "Failed to create User")
		return nil, e.NewInternalServerException(wrap)
	}

	authIdentity := &AuthIdentity{
		Provider:          ProviderPassword,
		UID:               uid,
		EncryptedPassword: digested,

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

func (r *repositoryImpl) Authenticate(uid string, password string) (*AuthClaim, e.Exception) {
	if strings.TrimSpace(uid) == "" || strings.TrimSpace(password) == "" {
		err := errors.New("Empty uid or password")
		return nil, e.NewUnauthorizedException(err)
	}

	aid := AuthIdentity{UID: uid}
	err := r.db.Where("uid = ?", uid).First(&aid).Error

	if err != nil {
		wrap := errors.Wrap(err, "Failed to find AuthIdentity with UID")

		if gorm.IsRecordNotFoundError(err) {
			return nil, e.NewUnauthorizedException(wrap)
		}

		return nil, e.NewInternalServerException(wrap)
	}

	if err := r.encryptor.Compare(aid.EncryptedPassword, password); err != nil {
		wrap := errors.Wrap(err, "Incorrect password")
		return nil, e.NewUnauthorizedException(wrap)
	}

	claim := aid.ToClaims()
	return claim, nil
}
