package main

import (
	"log"

	"github.com/labstack/echo/v4"

	"github.com/juri200405/uma-repo/app/internal/server/router"
	"github.com/juri200405/uma-repo/app/internal/infrastructures/dao/mysql"
)

func main() {
	db := mysql.ConnectDb()
	if err := mysql.MigrateDb(db); err != nil {
		log.Fatal(err)
	}
	e := echo.New()
	router.Router(e)
	e.Logger.Fatal(e.Start(":3000"))
}
