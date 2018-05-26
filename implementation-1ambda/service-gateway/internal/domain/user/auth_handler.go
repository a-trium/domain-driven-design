package user

import (
	"strings"

	e "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/exception"
	"github.com/pkg/errors"
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/pkg/generated/swagger/swagserver/swagapi"
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/pkg/generated/swagger/swagserver/swagapi/auth"
	"github.com/go-openapi/runtime/middleware"
	dto "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/pkg/generated/swagger/swagmodel"
	"strconv"
)

type AuthHandler interface {
	Configure(handlerRegistry *swagapi.GatewayAPI)
	Register(uid string, email string, password string) (*AuthClaim, e.Exception)
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

func (c *authHandlerImpl) Configure(registry *swagapi.GatewayAPI) () {
	registry.AuthRegisterHandler = auth.RegisterHandlerFunc(
		func(params auth.RegisterParams) middleware.Responder {
			if params.Body == nil {
				err := errors.New("Empty Body")
				ex := e.NewBadRequestException(err)
				return auth.NewLoginDefault(ex.StatusCode()).WithPayload(ex.ToSwaggerError())
			}

			uid := params.Body.UID
			email := params.Body.Email
			password := params.Body.Password

			claim, ex := c.Register(uid, email, password)
			if ex != nil {
				return auth.NewRegisterDefault(ex.StatusCode()).WithPayload(ex.ToSwaggerError())
			}

			response := dto.RegisterResponse{
				UID:    claim.UID,
				UserID: strconv.FormatUint(uint64(claim.UserID), 10),
			}
			return auth.NewRegisterOK().WithPayload(&response)
		})

	registry.AuthLoginHandler = auth.LoginHandlerFunc(
		func(params auth.LoginParams) middleware.Responder {
			if params.Body == nil {
				err := errors.New("Empty Body")
				ex := e.NewBadRequestException(err)
				return auth.NewLoginDefault(ex.StatusCode()).WithPayload(ex.ToSwaggerError())
			}

			uid := params.Body.UID
			password := params.Body.Password

			_, ex := c.Login(uid, password)
			if ex != nil {
				return auth.NewLoginDefault(ex.StatusCode()).WithPayload(ex.ToSwaggerError())
			}

			return auth.NewLoginOK().WithPayload(nil)
		})

	registry.AuthLogoutHandler = auth.LogoutHandlerFunc(
		func(params auth.LogoutParams) middleware.Responder {
			return auth.NewLogoutOK()
		})
}

func (c *authHandlerImpl) Register(uid string, email string, password string) (*AuthClaim, e.Exception) {
	if strings.TrimSpace(uid) == "" || strings.TrimSpace(password) == "" {
		err := errors.New("Empty uid or password")
		return nil, e.NewBadRequestException(err)
	}

	encrypted, err := c.encryptor.Digest(password)
	if err != nil {
		wrap := errors.Wrap(err, "Failed to digest password")
		return nil, e.NewInternalServerException(wrap)
	}

	aid, ex := c.userRepository.CreateAuthIdentity(uid, email, encrypted)
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
