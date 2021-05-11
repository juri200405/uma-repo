package router

import (
	"github.com/labstack/echo/v4"

	"github.com/juri200405/uma-repo/app/internal/server/handlers"
)

func Hello(e *echo.Echo) {
	e.GET("/", handlers.Hello())
}
