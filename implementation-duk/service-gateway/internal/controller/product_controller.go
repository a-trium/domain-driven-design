package controller

import (
	"fmt"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/product"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductController struct {
	repository product.Repository
}

func NewProductController(repository product.Repository) *ProductController {
	return &ProductController{repository}
}

func (ctrl *ProductController) GetProduct(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	product, err := ctrl.repository.FindById(id)

	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, product)
}