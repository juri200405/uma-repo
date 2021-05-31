package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/juri200405/uma-repo/app/internal/domain/models"
	"github.com/juri200405/uma-repo/app/internal/registry"
)

func RaceRegister(r *registry.RaceRegistry) echo.HandlerFunc {
	uc := r.GetRaceUsecase()
	return func(c echo.Context) error {
		u := &models.Race{}
		if err := c.Bind(u); err != nil {
			return err
		}
		fmt.Println(u)

		if err := uc.Register(u); err != nil {
			return err
		} else {
			return c.Redirect(http.StatusSeeOther, "/races")
		}
	}
}

func RacePage(r *registry.RaceRegistry) echo.HandlerFunc {
	uc := r.GetRaceUsecase()
	return func(c echo.Context) error {
		if races, err := uc.GetRaces(); err != nil {
			return err
		} else {
			return c.Render(http.StatusOK, "racePage", map[string]interface{}{"raceList": races})
		}
	}
}

func RaceDelete(r *registry.RaceRegistry) echo.HandlerFunc {
	uc := r.GetRaceUsecase()
	return func(c echo.Context) error {
		id, err := strconv.ParseUint(c.Param("raceID"), 10, 64)
		if err != nil {
			return err
		}
		if err := uc.Delete(uint(id)); err != nil {
			return err
		} else {
			return c.NoContent(http.StatusOK)
		}
	}
}