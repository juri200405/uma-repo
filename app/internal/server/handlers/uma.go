package handlers

import (
	"fmt"
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

		fmt.Printf("%#v\n", u)
		if err := uc.Register(u); err != nil {
			return err
		} else {
			return c.Redirect(http.StatusSeeOther, "/uma")
		}
	}
}

func UmaRegisterPage(r *registry.UmaRegistry) echo.HandlerFunc {
	uc := r.GetUmaUsecase()
	return func(c echo.Context) error {
		umas, err := uc.GetAll()
		if err != nil {
			return err
		}
		names, err := uc.GetNames()
		if err != nil {
			return err
		}
		factorList, err := uc.GetFactorList()
		if err != nil {
			return err
		}
		return c.Render(
			http.StatusOK,
			"umaRegisterPage",
			map[string]interface{}{
				"nameList": names,
				"umaList": umas,
				"factorList": factorList,
			},
		)
	}
}
