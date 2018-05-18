package main

import (
	"fmt"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/infra"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/web/config"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/web/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	env := config.GetEnvironment()
	//_ := config.GetLogger(env)
	gorm := config.GetDatabase(env.DatabaseProperty)

	userRepository := repository.NewUserRepository(gorm)
	userController := controller.NewUserController(userRepository)

	//user := domain.User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	//db.Create(&user)

	// TODO : middleware / router refactoring
	router.Use(dummyMiddleware)
	router.GET("/ping", controller.HealthCheck)

	v1 := router.Group("/v1")
	{
		v1.GET("/user/:userId", userController.User)
	}

	router.Run(":" + env.Port)
}

func dummyMiddleware(context *gin.Context) {
	fmt.Println("Dummy Middleware start")
	context.Next()
}
