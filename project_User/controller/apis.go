package controller

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// add middleware basic plus customizable
func Start(e *echo.Echo) {
	//starts an http server

	e.Use(middleware.BasicAuth(func(username string, password string, e echo.Context) (bool, error) {
		if username == "Atharv" && password == "lsq123" {
			return true, nil
		}
		return false, errors.New("error")
	}))
	e.GET("/user", GetUser)
	e.GET("/users", getAllUser)
	e.POST("/create", CreateUser)
	e.PUT("/user", UpdateUser)
}
