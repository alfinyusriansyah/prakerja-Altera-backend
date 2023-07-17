package routes

import (
	"uji-keterampilan-golang/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) *echo.Echo {

	// // Route untuk registrasi pengguna
	// e.POST("/register", controllers.Register)

	// // Route untuk login
	// e.POST("/login", controllers.Login)

	// api := e.Group("")
	// api.Use(middleware.JWT([]byte(os.Getenv("KEY_ENCRYPTOPN"))))

	e.POST("/books", controllers.CreateBook)
	e.GET("/books", controllers.GetAllBooks)
	e.GET("/authors/:AuthorId/books", controllers.GetBooksByAuthorId)
	e.DELETE("/books/:id", controllers.DeleteBook)
	e.PUT("/books/:id", controllers.UpdateBook)

	e.POST("/authors", controllers.CreateAuthor)
	e.GET("/authors", controllers.GetAllAuthor)
	e.DELETE("/authors/:id", controllers.DeleteAuthor)
	e.PUT("/authors/:id", controllers.UpdateAuthor)

	return e
}
