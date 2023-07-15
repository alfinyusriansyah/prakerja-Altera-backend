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

	if author.Name == "" {
		// Tangani jika nama author kosong atau buka string
		return c.JSON(http.StatusBadRequest, "Nama author tidak boleh kosong dan harus integer")
	}
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

func GetAllAuthor(c echo.Context) error {
	authors := []models.Author{}
	result := config.DB.Find(&authors)
	if result.Error != nil {
		// Tangani jika terjadi kesalahan saat mengambil buku
		return c.JSON(http.StatusInternalServerError, "Gagal mengambil buku")
	}
	return c.JSON(http.StatusOK, authors)
}
