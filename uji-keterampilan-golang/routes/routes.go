package routes

import (
	"uji-keterampilan-golang/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) *echo.Echo {
	e.POST("/books", controllers.CreateBook)
	e.GET("/books", controllers.GetAllBooks)
	e.DELETE("/books/:id", controllers.DeleteBook)

	e.POST("/authors", controllers.CreateAuthor)
	e.GET("/authors", controllers.GetAllAuthor)
	e.GET("/authors/:AuthorId/books", controllers.GetBooksByAuthorId)
	return e
}
