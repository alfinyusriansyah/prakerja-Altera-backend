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

func UpdateAuthor(c echo.Context) error {
	author_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: "Invalid author ID", Status: false, Data: nil,
		})
	}

	// Binding data dari body request ke objek buku
	var input models.Author
	if err := c.Bind(&input); err != nil {
		return err
	}

	result := config.DB.Model(&models.Author{}).Where("id = ?", author_id).Updates(input)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Failed to update Author", Status: false, Data: nil,
		})
	}

	// Periksa apakah ada data buku yang diupdate
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Author not found", Status: false, Data: nil,
		})
	}
	// Mengambil data buku yang telah diperbarui dari database
	var updatedAuthor models.Author
	if err := config.DB.First(&updatedAuthor, author_id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Failed to retrieve updated author", Status: false, Data: nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "author updated successfully",
		"author":  updatedAuthor,
	})
}
