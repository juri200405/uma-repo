package router

import (
	"github.com/labstack/echo/v4"

	"github.com/juri200405/uma-repo/app/internal/registry"
	"github.com/juri200405/uma-repo/app/internal/server/handlers"
)

func FactorRouter(e *echo.Echo) {
	r := registry.NewFactorRegistry()

	Route := e.Group("/factors")
	{
		Route.GET("", handlers.FactorPage(r))
		Route.POST("", handlers.FactorRegister(r))
		Route.DELETE("/:factorID", handlers.FactorDelete(r))
		Route.PATCH("/:factorID", handlers.FactorUpdate(r))
	}
}
