package main

import (
	"uji-keterampilan-golang/config"
	"uji-keterampilan-golang/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDatabase()
	e := echo.New()
	e = routes.InitRoute(e)
	e.Start(":8000")

}
