package controllers

import (
	"net/http"
	"uji-keterampilan-golang/config"
	"uji-keterampilan-golang/models"

	"github.com/labstack/echo/v4"
)

func CreateAuthor(c echo.Context) error {
	author := models.Author{}
	c.Bind(&author)

	result := config.DB.Create(&author)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal memasukkan data", Status: false, Data: nil,
		})
	}
	return c.JSON(http.StatusCreated, models.BaseResponse{
		Message: "Success memasukkan data", Status: true, Data: author,
	})
}

func GetBooksByAuthorId(c echo.Context) error {
	AuthorId := c.Param("AuthorId")

	author := models.Author{}
	result := config.DB.Preload("Books").First(&author, AuthorId)
	if result.Error != nil {
		// Tangani jika tidak dapat menemukan penulis atau kesalahan lainnya
		return c.JSON(http.StatusInternalServerError, "Gagal mendapatkan buku penulis")
	}

	return c.JSON(http.StatusOK, author.Books)
}
