package main

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/config"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/infra"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/web/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	gin := gin.Default()

	connection := config.GetDatabase(config.GetEnvironment())

	userRepository :=repository.NewUserRepository(connection.GetDB())
	userController := controller.NewUserController(userRepository)
	// Middleware

	// Routes
	gin.GET("/ping", controller.HealthCheck)

	groupV1 := gin.Group("/v1")
	groupV1.GET("/users/:id", userController.GetUser)
	groupV1.POST("/users", userController.AddUser)

	// Start server
	gin.Run(":8080")
}

