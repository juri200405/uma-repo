package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"

	"github.com/juri200405/uma-repo/app/internal/server/router"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	t := &Template{templates: template.Must(template.New("").ParseGlob("web/templates/**/*.html"))}
	e.Renderer = t
	e.Static("/public/css/", "./public/css/")
	e.Static("/public/js/", "./public/js/")

	router.Router(e)
	e.Logger.Fatal(e.Start(":3000"))
}
