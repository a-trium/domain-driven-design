package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func main() {

	// errors
	fmt.Println("hello")
	err := errors.New("Error~!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println(err)

	// gin
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

