package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Serve() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health-check", heathCheck)
	e.POST("/users/register", register)

	e.Logger.Fatal(e.Start(":5555"))
}
