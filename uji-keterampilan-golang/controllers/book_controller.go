package controllers

import (
	"net/http"
	"strconv"
	"uji-keterampilan-golang/config"
	"uji-keterampilan-golang/models"

	"github.com/labstack/echo/v4"
)

// membuta dan menyimpan buku
func CreateBook(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	if book.Name == "" {
		// Tangani jika nama book kosong atau
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: "nama harus diisi", Status: false, Data: nil,
		})
	}
	// Validasi field year
	if book.Year <= 0 {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: "Tahun harus berupa bilangan bulat positif", Status: false, Data: nil,
		})
	}
	// Validasi field PageCount
	if book.PageCount <= 0 {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: "Jumlah halaman harus berupa bilangan bulat positif", Status: false, Data: nil,
		})
	}
	// configurasi DB
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

// Menampilkan Buku bedasarkan id author
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

func DeleteBook(c echo.Context) error {
	book_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid book ID",
		})
	}

	// Lanjutkan proses penghapusan data buku dari database
	book := models.Book{}
	result := config.DB.Delete(book, book_id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Failed to delete book", Status: false, Data: nil,
		})
	}

	// Periksa apakah ada data buku yang dihapus
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Message: "Book not found", Status: false, Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Book deleted successfully",
	})
}
