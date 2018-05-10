package rest

import (
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/pkg/generated/swagger/swagserver/swagapi"
	"go.uber.org/zap"
)

type Controller interface {
	Configure(api *swagapi.GatewayAPI)
}

type controllerImpl struct {
	logger *zap.SugaredLogger
}

func NewController(logger *zap.SugaredLogger) Controller {
	return &controllerImpl{logger: logger,}
}

func (ctrl *controllerImpl) Configure(api *swagapi.GatewayAPI) {
}
