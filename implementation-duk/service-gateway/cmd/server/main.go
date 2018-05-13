package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message" : "pong",
		})
	})

	r.Run()	// listen and serve on 0.0.0.0:8080
}
