package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//db := config.GetDatabase()
	//user := domain.User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	//db.Create(&user)

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message" : "pong",
		})
	})

	r.Run()	// listen and serve on 0.0.0.0:8080
}