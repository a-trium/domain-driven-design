package controller

import (
	"fmt"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CustomerController struct {
	repository user.Repository
}

func NewCustomerController(repository user.Repository) *CustomerController {
	return &CustomerController{repository:repository}
}

func (ctrl *CustomerController) GetCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, ctrl.repository.FindOne(id))
}

func (ctrl *CustomerController) AddCustomer(c *gin.Context) {
	customer := new(user.Customer)
	if err := c.Bind(customer); err != nil {
		fmt.Println(err.Error())
	}
	ctrl.repository.Save(customer)

	c.JSON(http.StatusOK, customer)
}