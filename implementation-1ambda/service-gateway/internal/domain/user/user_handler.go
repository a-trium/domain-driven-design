package user

import (
	"strings"

	e "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/exception"
	dto "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/pkg/generated/swagger/swagmodel"
	api "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/pkg/generated/swagger/swagserver/swagapi/auth"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type AuthHandler interface {
	Register(params api.RegisterParams) (*dto.Empty, *dto.Error)
	Login(params api.LoginParams) (*AuthClaim, *dto.Error)
	Logout(params api.LogoutParams) (*dto.Empty, *dto.Error)
}

type authHandlerImpl struct {
	logger         *zap.SugaredLogger
	userRepository Repository
	encryptor      Encryptor
}

func NewAuthHandler(logger *zap.SugaredLogger, repo Repository, encryptor Encryptor) AuthHandler {
	return &authHandlerImpl{
		logger:         logger,
		userRepository: repo,
		encryptor:      encryptor,
	}
}

func (c *authHandlerImpl) Register(params api.RegisterParams) (*dto.Empty, *dto.Error) {
	uid := params.Body.UID
	password := params.Body.Password

	// TODO: logging

	if strings.TrimSpace(uid) == "" || strings.TrimSpace(password) == "" {
		err := errors.New("Empty uid or password")
		ex := e.NewBadRequestException(err)
		return nil, ex.ToSwaggerError()
	}

	encrypted, err := c.encryptor.Digest(password)
	if err != nil {
		wrap := errors.Wrap(err, "Failed to digest password")
		ex := e.NewInternalServerException(wrap)
		return nil, ex.ToSwaggerError()
	}

	_, ex := c.userRepository.CreateAuthIdentity(uid, encrypted)
	if ex != nil {
		return nil, ex.ToSwaggerError()
	}

	return nil, nil
}

func (c *authHandlerImpl) Login(params api.LoginParams) (*AuthClaim, *dto.Error) {
	uid := params.Body.UID
	password := params.Body.Password

	if strings.TrimSpace(uid) == "" || strings.TrimSpace(password) == "" {
		err := errors.New("Empty uid or password")
		ex := e.NewUnauthorizedException(err)
		return nil, ex.ToSwaggerError()
	}

	aid, ex := c.userRepository.FindAuthIdentityByUID(uid)
	if ex != nil {
		return nil, ex.ToSwaggerError()
	}

	if err := c.encryptor.Compare(aid.EncryptedPassword, password); err != nil {
		wrap := errors.Wrap(err, "Incorrect password")
		ex := e.NewUnauthorizedException(wrap)
		return nil, ex.ToSwaggerError()
	}

	return aid.ToClaims(), nil
}

func (c *authHandlerImpl) Logout(params api.LogoutParams) (*dto.Empty, *dto.Error) {
	return nil, nil
}
