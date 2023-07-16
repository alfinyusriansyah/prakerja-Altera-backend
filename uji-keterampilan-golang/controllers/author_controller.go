package controllers

import (
	"net/http"
	"strconv"
	"uji-keterampilan-golang/config"
	"uji-keterampilan-golang/models"

	"github.com/labstack/echo/v4"
)

// ---------------------- Create Author controller ------------------
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

// ---------------------- Get All Author controller ------------------
func GetAllAuthor(c echo.Context) error {
	authors := []models.Author{}
	result := config.DB.Find(&authors)
	if result.Error != nil {
		// Tangani jika terjadi kesalahan saat mengambil buku
		return c.JSON(http.StatusInternalServerError, "Gagal mengambil buku")
	}
	return c.JSON(http.StatusOK, authors)
}

// ---------------------- Delete Author controller ------------------
func DeleteAuthor(c echo.Context) error {
	author_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid author ID",
		})
	}
	// Lanjutkan proses penghapusan data author dari database
	author := models.Author{}
	result := config.DB.Delete(author, author_id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Failed to delete Author", Status: false, Data: nil,
		})
	}
	// Periksa apakah ada data auhtor yang dihapus
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Message: "Id Author not found", Status: false, Data: nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Author deleted successfully",
	})
}
