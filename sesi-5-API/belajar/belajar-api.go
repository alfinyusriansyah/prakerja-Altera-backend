package main

import (
	"github.com/labstack/echo/v4"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	e := echo.New()

	e.GET("/helloworld", HelloController)
	e.GET("/users", UserController)
	e.POST("/users", CreateUser)
	e.Start(":8000")
}

func CreateUser(c echo.Context) error {
	user := User{}
	c.Bind(&user)
	return c.JSON(200, user)
}

func UserController(c echo.Context) error {
	var user []User
	user = append(user, User{1, "alfin", "alfin@gmail.com"})
	user = append(user, User{2, "yusri", "yusri@gmail.com"})

	return c.JSON(200, user)
}

func HelloController(c echo.Context) error {
	return c.JSON(200, "HELLOWORLD")
}
