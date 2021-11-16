package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/juri200405/uma-repo/app/internal/domain/models"
	"github.com/juri200405/uma-repo/app/internal/registry"
)

type (
	RaceResultList []models.RaceResult
	WhiteFactors   []uint
)

func (r *RaceResultList) UnmarshalParam(param string) error {
	return json.Unmarshal([]byte(param), r)
}
func (w *WhiteFactors) UnmarshalParam(param string) error {
	return json.Unmarshal([]byte(param), w)
}

func UmaRegister(r *registry.UmaRegistry) echo.HandlerFunc {
	uc := r.GetUmaUsecase()
	return func(c echo.Context) error {
		u := &models.Uma{}
		if err := c.Bind(u); err != nil {
			return err
		}
		fmt.Printf("%#v\n", u)

		var wFactorIds WhiteFactors
		var r RaceResultList
		if err := echo.FormFieldBinder(c).
			BindUnmarshaler("white_factor_items", &wFactorIds).
			BindUnmarshaler("race_results", &r).
			BindError(); err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(wFactorIds)
		fmt.Println(r)

		if err := uc.Register(u, wFactorIds, r); err != nil {
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
		races, err := uc.GetRaces("turn, place, ground, length")
		if err != nil {
			return err
		}
		return c.Render(
			http.StatusOK,
			"umaDetailPage",
			map[string]interface{}{
				"purpose":         "登録",
				"action":          "/uma",
				"nameList":        names,
				"umaList":         umas,
				"raceList":        races,
				"blueFactorList":  blueFactorList,
				"redFactorList":   redFactorList,
				"whiteFactorList": whiteFactorList,
			},
		)
	}
}

func UmaListPage(r *registry.UmaRegistry) echo.HandlerFunc {
	uc := r.GetUmaUsecase()
	return func(c echo.Context) error {
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

func UmaUpdater(r *registry.UmaRegistry) echo.HandlerFunc {
	uc := r.GetUmaUsecase()
	return func(c echo.Context) error {
		var umaID uint
		if err := echo.PathParamsBinder(c).Uint("umaID", &umaID).BindError(); err != nil {
			return err
		}
		u, err := uc.FindById(umaID)
		if err != nil {
			return err
		}
		u.Transferred = false
		if err := c.Bind(&u); err != nil {
			return err
		}
		fmt.Printf("%#v\n", u)

		var wFactorIds WhiteFactors
		var r RaceResultList
		if err := echo.FormFieldBinder(c).
			BindUnmarshaler("white_factor_items", &wFactorIds).
			BindUnmarshaler("race_results", &r).
			BindError(); err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(wFactorIds)
		fmt.Println(r)

		if err := uc.Update(&u, wFactorIds, r); err != nil {
			return err
		} else {
			return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/uma/%d", umaID))
		}
	}
}

func UmaDetailPage(r *registry.UmaRegistry) echo.HandlerFunc {
	uc := r.GetUmaUsecase()
	return func(c echo.Context) error {
		var umaID uint
		if err := echo.PathParamsBinder(c).Uint("umaID", &umaID).BindError(); err != nil {
			return err
		}
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
		races, err := uc.GetRaces("turn, place, ground, length")
		if err != nil {
			return err
		}
		uma, err := uc.FindById(umaID)
		if err != nil {
			return err
		}
		return c.Render(
			http.StatusOK,
			"umaDetailPage",
			map[string]interface{}{
				"purpose":         "編集",
				"action":          fmt.Sprintf("/uma/%d", umaID),
				"nameList":        names,
				"umaList":         umas,
				"raceList":        races,
				"blueFactorList":  blueFactorList,
				"redFactorList":   redFactorList,
				"whiteFactorList": whiteFactorList,
				"Uma":             uma,
			},
		)
	}
}

func GetAllUmas(r *registry.UmaRegistry) echo.HandlerFunc {
	uc := r.GetUmaUsecase()
	return func(c echo.Context) error {
		umas, err := uc.GetAll()
		if err != nil {
			fmt.Println(err)
			return err
		}
		return c.JSON(
			http.StatusOK,
			umas,
		)
	}
}
