package controller

import (
	"fmt"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type userController struct {
	repository user.Repository
}

func NewUserController(repository user.Repository) *userController {
	return &userController{repository:repository}
}

func (ctrl *userController) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, ctrl.repository.FindOne(id))
}

func (ctrl *userController) AddUser(c *gin.Context) {
	user := new(user.User)
	if err := c.Bind(user); err != nil {
		fmt.Println(err.Error())
	}
	ctrl.repository.Save(user)

	c.JSON(http.StatusOK, user)
}