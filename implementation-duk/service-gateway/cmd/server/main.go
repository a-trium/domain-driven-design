package main

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/web/controller"
	"github.com/labstack/echo"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	// Routes
	e.GET("/ping", controller.HealthCheck)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

