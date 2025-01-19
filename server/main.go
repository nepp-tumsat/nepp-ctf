package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Welcome NePP CTF!!")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Flag string }{Flag: "nepp{network-http-q1-883c3944-9361-4db8-8364-69db318c14b7}"})
	})

	httpPort := "8080"
	e.Logger.Fatal(e.Start(":" + httpPort))
}