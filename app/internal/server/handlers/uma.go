package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/labstack/echo/v4"

	"github.com/juri200405/uma-repo/app/internal/registry"
	"github.com/juri200405/uma-repo/app/internal/domain/models"
)

type RaceResultList []models.RaceResult
func (r *RaceResultList) UnmarshalParam(param string) error {
	return json.Unmarshal([]byte(param), r)
}

func UmaRegister(r *registry.UmaRegistry) echo.HandlerFunc {
	uc := r.GetUmaUsecase()
	return func(c echo.Context) error {
		u := &models.Uma{}
		if err := c.Bind(u); err != nil {
			return err
		}
		fmt.Printf("%#v\n", u)

		var wFactorIds []uint
		var r RaceResultList
		if err := echo.FormFieldBinder(c).
			BindWithDelimiter("white_factor_items", &wFactorIds, ",").
			BindUnmarshaler("race_results", &r).
			BindError(); err != nil {
				fmt.Println(err)
				return err
			}
		fmt.Println(wFactorIds)
		fmt.Println(r)

		if err := uc.Register(u, wFactorIds[1:], r); err != nil {
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
		blueFactorList, redFactorList, whiteFactorList, err := uc.GetFactorList()
		if err != nil {
			return err
		}
		races, err := uc.GetRaces()
		if err != nil {
			return err
		}
		return c.Render(
			http.StatusOK,
			"umaRegisterPage",
			map[string]interface{}{
				"nameList": names,
				"umaList": umas,
				"raceList": races,
				"blueFactorList": blueFactorList,
				"redFactorList": redFactorList,
				"whiteFactorList": whiteFactorList,
			},
		)
	}
}

func UmaListPage(r *registry.UmaRegistry) echo.HandlerFunc {
	uc := r.GetUmaUsecase()
	return func(c echo.Context) error {
		fmt.Println("uma list page")
		umas, err := uc.GetAll()
		if err != nil {
			fmt.Println(err)
			return err
		}
		return c.Render(
			http.StatusOK,
			"umaPage",
			map[string]interface{}{
				"umaList": umas,
			},
		)
	}
}
