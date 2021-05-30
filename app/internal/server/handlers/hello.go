package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Race struct {
	Name   string `json:name`
	Result uint   `json:result`
}
type Races []Race

func (r *Races) UnmarshalParam(param string) error {
	return json.Unmarshal([]byte(param), r)
}

func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		v := c.Param("var")
		return c.Render(http.StatusOK, "helloPage", map[string]string{"Word": v})
	}
}

func GoodDay() echo.HandlerFunc {
	return func(c echo.Context) error {
		v := c.Param("var")
		return c.Render(http.StatusOK, "goodDayPage", map[string]string{"Word": v})
	}
}

func Object() echo.HandlerFunc {
	return func(c echo.Context) error {
		var r Races
		if err := echo.FormFieldBinder(c).BindUnmarshaler("list", &r).BindError(); err != nil {
			fmt.Println(err)
			return c.String(http.StatusOK, "エラーだよ")
		}
		fmt.Println(r)
		return c.String(http.StatusOK, "エラーじゃないよ")
	}
}
