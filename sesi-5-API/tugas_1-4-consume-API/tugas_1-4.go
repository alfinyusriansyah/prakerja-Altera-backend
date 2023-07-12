// Maaf Sebelumnya sebetulnya saya masih agak binguang
// dengan penggunaan dan cara kerja method - method yang ada di codingan bawah ini.
// Mohon bisa di jelaskan

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Post struct {
	UserId int    `json:"userid"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	e := echo.New()

	e.GET("/posts", GetPost)
	e.GET("/posts/:id", DetailPost)
	e.POST("/posts", SavePost)
	e.DELETE("/posts/:id", DeletePost)
	e.Start(":8000")
}

// ----------------- Mendapatkan data --------------------
func GetPost(c echo.Context) error {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Membaca respons JSON
	var posts []Post
	err = json.NewDecoder(response.Body).Decode(&posts)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, posts)

}

// ---------------- Menampilkan data berdaasarkan ID -----------------
func DetailPost(c echo.Context) error {
	var id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Membaca respons JSON
	var posts []Post
	err = json.NewDecoder(response.Body).Decode(&posts)
	if err != nil {
		log.Fatal(err)
	}

	// perulangan mencari id yang sesuai request, di struct Post[]
	var postWithId Post
	for _, post := range posts {
		if id == post.Id {
			postWithId = post
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success display post by id",
				"post":     postWithId,
			})

		}
	}
	return c.JSON(echo.ErrBadRequest.Code, map[string]interface{}{
		"messages": "Post Not Found",
		"post":     nil,
	})
}

// ----------------------- Meyimpan Data -------------------------
func SavePost(c echo.Context) error {
	post := new(Post)
	if err := c.Bind(post); err != nil {
		return err
	}

	// Mengubah data menjadi format JSON
	jsonData, err := json.Marshal(post)
	if err != nil {
		log.Fatal(err)
	}

	// Membuat permintaan POST ke API
	response, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusCreated {
		log.Println("Postingan berhasil disimpan.")
	} else {
		log.Println("Gagal menyimpan postingan.")
	}

	// Mengembalikan respons berupa data postingan
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create post",
		"post":     post,
	})
}

// ---------------- Hapus Data -------------------
func DeletePost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	// Membuat permintaan DELETE ke API
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Mengirim permintaan DELETE ke API
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		log.Println("Postingan berhasil dihapus.")
	} else {
		log.Println("Gagal menghapus postingan.")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete post",
	})
}
