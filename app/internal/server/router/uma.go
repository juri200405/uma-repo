package router

import (
	"github.com/labstack/echo/v4"

	"github.com/juri200405/uma-repo/app/internal/registry"
	"github.com/juri200405/uma-repo/app/internal/server/handlers"
)

func UmaRouter(e *echo.Echo) {
	r := registry.NewUmaRegistry()

	e.GET("/", handlers.UmaListPage(r))
	Route := e.Group("/uma")
	{
		Route.GET("", handlers.UmaRegisterPage(r))
		Route.POST("", handlers.UmaRegister(r))
		Route.GET("/:umaID", handlers.UmaDetailPage(r))
		Route.POST("/:umaID", handlers.UmaUpdater(r))
		Route.GET("/json", handlers.GetAllUmas(r))
	}
}
