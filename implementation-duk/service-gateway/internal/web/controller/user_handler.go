package controller

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/user"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type userController struct {
	repository user.Repository
}

func NewUserController(repository user.Repository) *userController {
	return &userController{repository:repository}
}

func (ctrl *userController) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, ctrl.repository.FindOne(id))
}

func (ctrl *userController) AddUser(c echo.Context) error {
	user := new(user.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	ctrl.repository.Save(user)

	return c.JSON(http.StatusOK, user)
}