package user

import (
	dto "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/pkg/generated/swagger/swagmodel"
	api "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/pkg/generated/swagger/swagserver/swagapi/auth"
	"go.uber.org/zap"
)

type UserHandler interface {
	Register(params api.RegisterParams) (*dto.Empty, *dto.Error)
	Login(params api.LoginParams) (*dto.Empty, *dto.Error)
	Logout(params api.LogoutParams) (*dto.Empty, *dto.Error)
}

type userHandlerImpl struct {
	logger         *zap.SugaredLogger
	userRepository Repository
}

func NewUserHandler(logger *zap.SugaredLogger, repo Repository) UserHandler {
	return &userHandlerImpl{logger: logger, userRepository: repo,}
}

func (c *userHandlerImpl) Register(params api.RegisterParams) (*dto.Empty, *dto.Error) {
	return nil, nil
}

func (c *userHandlerImpl) Login(params api.LoginParams) (*dto.Empty, *dto.Error) {
	return nil, nil
}

func (c *userHandlerImpl) Logout(params api.LogoutParams) (*dto.Empty, *dto.Error) {
	return nil, nil
}
