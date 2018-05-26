package user

import (
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/pkg/generated/swagger/swagserver/swagapi"
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/pkg/generated/swagger/swagserver/swagapi/auth"
	"github.com/go-openapi/runtime/middleware"
)

func RegisterAuthHandler(register *swagapi.GatewayAPI, authHandler AuthHandler) {

	register.AuthLoginHandler = auth.LoginHandlerFunc(
		func(params auth.LoginParams) middleware.Responder {
			uid := params.Body.UID
			password := params.Body.Password

			_, ex := authHandler.Login(uid, password)
			if ex != nil {
				return auth.NewLoginDefault(ex.StatusCode()).WithPayload(ex.ToSwaggerError())
			}

			return auth.NewLoginOK().WithPayload(nil)
		})

}
