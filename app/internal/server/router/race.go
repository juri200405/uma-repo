package router

import (
	"github.com/labstack/echo/v4"

	"github.com/juri200405/uma-repo/app/internal/registry"
	"github.com/juri200405/uma-repo/app/internal/server/handlers"
)

func RaceRouter(e *echo.Echo) {
	r := registry.NewRaceRegistry()

	Route := e.Group("/races")
	{
		Route.GET("", handlers.RacePage(r))
		Route.POST("", handlers.RaceRegister(r))
	}
}
