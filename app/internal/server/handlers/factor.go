package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/juri200405/uma-repo/app/internal/registry"
	"github.com/juri200405/uma-repo/app/internal/domain/models"
)

func FactorRegister(r *registry.FactorRegistry) echo.HandlerFunc {
	uc := r.GetFactorUsecase()
	return func(c echo.Context) error {
		u := &models.Factor{}
		if err := c.Bind(u); err != nil {
			return err
		}
		fmt.Println(u)

		if err := uc.Register(u); err != nil {
			return err
		} else {
			return c.Redirect(http.StatusSeeOther, "/factors")
		}
	}
}

func FactorPage(r *registry.FactorRegistry) echo.HandlerFunc {
	uc := r.GetFactorUsecase()
	return func(c echo.Context) error {
		if factors, err := uc.GetAll(); err != nil {
			return err
		} else {
			return c.Render(http.StatusOK, "factorPage", map[string]interface{}{"factorList":factors})
		}
	}
}
