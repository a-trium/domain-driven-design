package main

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/config"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/infra"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/web/controller"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func main() {

	route := gin.Default()

	c := dig.New()
	c.Provide(db)
	c.Provide(repository.NewUserRepository)
	c.Provide(controller.NewUserController)

	c.Invoke(func(ctrl *controller.UserController) {
		groupV1 := route.Group("/v1")
		groupV1.GET("/users/:id", ctrl.GetUser)
		groupV1.POST("/users", ctrl.AddUser)
	})
	c.Invoke(func() {
		route.GET("/ping", controller.HealthCheck)
	})

	route.Run(":8080")
}

func db() *config.DBConnection {
	return config.GetDatabase(config.GetEnvironment())
}
