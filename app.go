package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func RootPage(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func main() {
	e := echo.New()
	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", RootPage)

	e.Logger.Fatal(e.Start(":1323"))
}
