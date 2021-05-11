package router

import (
	"github.com/labstack/echo/v4"

	"github.com/juri200405/uma-repo/app/internal/registry"
	"github.com/juri200405/uma-repo/app/internal/server/handlers"
)

func UmaRouter(e *echo.Echo) {
	r := registry.NewUmaRegistry()

	Route := e.Group("/uma")
	{
		Route.POST("", handlers.UmaRegister(r))
	}
}
