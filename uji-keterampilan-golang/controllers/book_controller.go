package controllers

import (
	"net/http"
	"uji-keterampilan-golang/config"
	"uji-keterampilan-golang/models"

	"github.com/labstack/echo/v4"
)

// membuta dan menyimpan buku
func CreateBook(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

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

// Menampilkan semua buku
func GetAllBooks(c echo.Context) error {
	books := []models.Book{}

	result := config.DB.Find(&books)
	if result.Error != nil {
		// Tangani jika terjadi kesalahan saat mengambil buku
		return c.JSON(http.StatusInternalServerError, "Gagal mengambil buku")
	}

	return c.JSON(http.StatusOK, books)
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
