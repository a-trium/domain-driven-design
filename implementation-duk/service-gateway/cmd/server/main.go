package main

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/config"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/controller"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/infra"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func main() {

	route := gin.Default()

	c := dig.New()
	c.Provide(config.GetEnvironment)
	c.Provide(config.GetDatabase)
	c.Provide(config.GetLogger)

	c.Provide(repository.NewCustomerRepository)
	c.Provide(repository.NewProductRepository)

	c.Provide(controller.NewCustomerController)
	c.Provide(controller.NewProductController)

	// health check
	c.Invoke(healthCheckHandler(route))

	// customer
	groupV1 := route.Group("/v1")
	c.Invoke(func(ctrl *controller.CustomerController) {
		groupV1.GET("/customers/:id", ctrl.GetCustomer)
		groupV1.POST("/customers", ctrl.AddCustomer)
	})

	c.Invoke(func(ctrl *controller.ProductController) {
		groupV1.GET("/products/:id", ctrl.GetProduct)
		groupV1.GET("/products", ctrl.GetProductsByTagId)
	})

	route.Run(":8080")
}

func healthCheckHandler(route *gin.Engine) func() {
	return func() {
		route.GET("/ping", controller.HealthCheck)
	}
}
