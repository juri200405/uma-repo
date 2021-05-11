package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/juri200405/uma-repo/app/internal/registry"
	"github.com/juri200405/uma-repo/app/internal/domain/models"
)

func UmaRegister(r *registry.UmaRegistry) echo.HandlerFunc {
	uc := r.GetUmaUsecase()
	return func(c echo.Context) error {
		u := &models.Uma{}
		if err := c.Bind(u); err != nil {
			return err
		}

		if err := uc.Register(u); err != nil {
			return err
		} else {
			return c.NoContent(http.StatusOK)
		}
	}
}
