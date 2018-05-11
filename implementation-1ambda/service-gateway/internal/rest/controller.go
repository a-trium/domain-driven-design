package rest

import (
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/pkg/generated/swagger/swagserver/swagapi"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type Controller interface {
	Configure(api *swagapi.GatewayAPI)
}

type controllerImpl struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewController(db *gorm.DB, logger *zap.SugaredLogger) Controller {
	return &controllerImpl{db: db, logger: logger,}
}

func (ctrl *controllerImpl) Configure(api *swagapi.GatewayAPI) {
}
