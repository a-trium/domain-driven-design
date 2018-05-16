package main

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	env := config.GetEnvironment()
	config.GetDatabase(env.DatabaseProperty)
	//user := domain.User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	//db.Create(&user)


	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message" : "pong",
		})
	})

	r.Run()	// listen and serve on 0.0.0.0:8080
}
