package httpserver

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"shopstoretest/cfg/hamedcfg"
	"shopstoretest/repository/mysql"
	"shopstoretest/service/userservice"
)

type Reader struct {
	echo *echo.Echo
}

func register(c echo.Context) error {
	req := userservice.RegisterRequest{}

	//fmt.Println("here")

	bErr := c.Bind(&req)
	if bErr != nil {

		return bErr
	}

	myRepo := mysql.New(hamedcfg.Database)

	//fmt.Println("hi i'm here")
	uService := userservice.Service{Repository: myRepo}
	res, err := uService.Register(req)
	if err != nil {

		return fmt.Errorf("cant register user %w", err)
	}

	return c.JSON(http.StatusOK, res)
}
