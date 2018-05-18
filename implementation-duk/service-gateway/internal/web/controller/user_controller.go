package controller

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/user"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct {
	repository user.Repository
}

func NewUserController(repository user.Repository) *UserController {
	return &UserController{repository: repository}
}

func (this *UserController) User(c *gin.Context) {

	userId := c.Param("userId")
	id, _ := strconv.Atoi(userId)
	c.JSON(200, gin.H{
		"user": this.repository.FindOne(id),
	})
}
