package routes

import (
	"uji-keterampilan-golang/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) *echo.Echo {
	e.POST("/book", controllers.CreateBook)
	return e
}
