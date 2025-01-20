package network_http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)


func Q2Handler(c echo.Context) (err error) {
	req := new(q2Request)
	if err = c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
    }

	if req.MeijiMaru == 5 && req.MarineCafe == 15 && req.KaioDormitory == 24 {
		return c.JSON(http.StatusOK, q2Response{Flag: "nepp{network-http-q2-f4972292-017c-4de5-be5d-b369750e0542}"})
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "The numbers on the map of the Etchujima Campus are incorrect."})
	}
}

type q2Request struct {
	MeijiMaru  int32 `json:"meiji_maru" validate:"required"`
	MarineCafe int32 `json:"marine_cafe" validate:"required"`
	KaioDormitory int32 `json:"kaio_dormitory" validate:"required"`
}

type q2Response struct {
	Flag  string `json:"flag"`
}