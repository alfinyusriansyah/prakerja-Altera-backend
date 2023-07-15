package main

import (
	"uji-keterampilan-golang/config"
	"uji-keterampilan-golang/routes"

	"github.com/labstack/echo/v4"
)

// fungsi main
func main() {
	// koneksi database
	config.ConnectDatabase()
	// memanggil library / importan echo framework
	e := echo.New()
	// memanggil route
	e = routes.InitRoute(e)
	// host
	e.Start(":8000")

}
