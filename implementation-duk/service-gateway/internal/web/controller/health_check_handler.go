package controller

import (
	"github.com/labstack/echo"
)

type healthCheck struct {
	Ping string `json:"ping" xml:"ping"`
}

func HealthCheck(c echo.Context) error {
	return c.JSON(200, healthCheck{"pong"})
}