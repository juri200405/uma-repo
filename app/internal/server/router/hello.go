package router

import (
	"github.com/labstack/echo/v4"

	"github.com/juri200405/uma-repo/app/internal/server/handlers"
)

func Hello(e *echo.Echo) {
	e.GET("/hello/:var", handlers.Hello())
	e.GET("/goodday/:var", handlers.GoodDay())
	e.POST("/obj", handlers.Object())
}
