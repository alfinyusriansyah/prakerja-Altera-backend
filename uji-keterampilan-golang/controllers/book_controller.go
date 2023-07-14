package controllers

import (
	"net/http"
	"time"
	"uji-keterampilan-golang/config"
	"uji-keterampilan-golang/models"

	"github.com/labstack/echo/v4"
)

func CreateBook(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	currentTime := time.Now()
	book.CreatedAt = currentTime
	book.UpdatedAt = currentTime

	result := config.DB.Create(&book)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal memasukkan data", Status: false, Data: nil,
		})
	}
	return c.JSON(http.StatusCreated, models.BaseResponse{
		Message: "Success memasukkan data", Status: true, Data: book,
	})
}
