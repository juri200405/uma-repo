package router

import (
	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	Hello(e)
	UmaRouter(e)
	FactorRouter(e)
}
