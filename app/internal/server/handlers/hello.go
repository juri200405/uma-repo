package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		v := c.Param("var")
		return c.Render(http.StatusOK, "helloPage", map[string]string{"Word": v})
	}
}

func GoodDay() echo.HandlerFunc {
	return func(c echo.Context) error {
		v := c.Param("var")
		return c.Render(http.StatusOK, "goodDayPage", map[string]string{"Word": v})
	}
}
