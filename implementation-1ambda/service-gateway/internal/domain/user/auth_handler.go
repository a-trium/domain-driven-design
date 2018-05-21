package user

import (
	"strings"

	e "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/exception"
	api "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/pkg/generated/swagger/swagserver/swagapi/auth"
	"github.com/pkg/errors"
)

type AuthHandler interface {
	Register(params api.RegisterParams) (e.Exception)
	Login(params api.LoginParams) (*AuthClaim, e.Exception)
	Logout(params api.LogoutParams) (e.Exception)
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

func (c *authHandlerImpl) Register(params api.RegisterParams) (e.Exception) {
	uid := params.Body.UID
	password := params.Body.Password

	// TODO: logging

	if strings.TrimSpace(uid) == "" || strings.TrimSpace(password) == "" {
		err := errors.New("Empty uid or password")
		return e.NewBadRequestException(err)
	}

	encrypted, err := c.encryptor.Digest(password)
	if err != nil {
		wrap := errors.Wrap(err, "Failed to digest password")
		return e.NewInternalServerException(wrap)
	}

	if _, ex := c.userRepository.CreateAuthIdentity(uid, encrypted); ex != nil {
		return ex
	}

	return nil
}

func (c *authHandlerImpl) Login(params api.LoginParams) (*AuthClaim, e.Exception) {
	uid := params.Body.UID
	password := params.Body.Password

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

func (c *authHandlerImpl) Logout(params api.LogoutParams) (e.Exception) {
	return nil
}
