package user

import (
	"strings"

	e "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/exception"
	"github.com/pkg/errors"
)

type AuthHandler interface {
	Register(uid string, password string) (*AuthClaim, e.Exception)
	Login(uid string, password string) (*AuthClaim, e.Exception)
	Logout() (e.Exception)
}

type authHandlerImpl struct {
	userRepository Repository
	encryptor      Encryptor
}

func NewAuthHandler(repo Repository, encryptor Encryptor) AuthHandler {
	return &authHandlerImpl{
		userRepository: repo,
		encryptor:      encryptor,
	}
}

func (c *authHandlerImpl) Register(uid string, password string) (*AuthClaim, e.Exception) {
	if strings.TrimSpace(uid) == "" || strings.TrimSpace(password) == "" {
		err := errors.New("Empty uid or password")
		return nil, e.NewBadRequestException(err)
	}

	encrypted, err := c.encryptor.Digest(password)
	if err != nil {
		wrap := errors.Wrap(err, "Failed to digest password")
		return nil, e.NewInternalServerException(wrap)
	}

	aid, ex := c.userRepository.CreateAuthIdentity(uid, encrypted)
	if ex != nil {
		return nil, ex
	}

	return aid.ToClaims(), nil
}

func (c *authHandlerImpl) Login(uid string, password string) (*AuthClaim, e.Exception) {
	if strings.TrimSpace(uid) == "" || strings.TrimSpace(password) == "" {
		err := errors.New("Empty uid or password")
		return nil, e.NewUnauthorizedException(err)
	}

	aid, ex := c.userRepository.FindAuthIdentityByUID(uid)
	if ex != nil {
		return nil, ex
	}

	if err := c.encryptor.Compare(aid.EncryptedPassword, password); err != nil {
		wrap := errors.Wrap(err, "Incorrect password")
		return nil, e.NewUnauthorizedException(wrap)
	}

	return aid.ToClaims(), nil
}

func (c *authHandlerImpl) Logout() (e.Exception) {
	return nil
}
