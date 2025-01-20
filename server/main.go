package main

import (
	"net/http"
	"server/network_http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Welcome NePP CTF!!!")
	})

	// For Network - HTTP Q1
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Flag string }{Flag: "nepp{network-http-q1-883c3944-9361-4db8-8364-69db318c14b7}"})
	})
	// For Network - HTTP Q2
	e.POST("/network-http-q2", func(c echo.Context) error {
		return network_http.Q2Handler(c)
	})

	httpPort := "50001"
	e.Logger.Fatal(e.Start(":" + httpPort))
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
