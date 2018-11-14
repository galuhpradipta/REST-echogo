package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", welcome)

	e.Logger.Fatal(e.Start(":5000"))
}

func welcome(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "welcome !!",
	})
}
