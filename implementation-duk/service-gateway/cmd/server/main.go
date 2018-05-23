package main

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/config"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/infra"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/web/controller"
	"github.com/labstack/echo"
)

func main() {

	// Echo instance
	e := echo.New()


	connection := config.GetDatabase(config.GetEnvironment())

	userRepository :=repository.NewUserRepository(connection.GetDB())

	// Middleware
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	userController := controller.NewUserController(userRepository)

	// Routes
	e.GET("/ping", controller.HealthCheck)

	groupV1 := e.Group("/v1")
	groupV1.GET("/users/:id", userController.GetUser)
	groupV1.POST("/users", userController.AddUser)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

