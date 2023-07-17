package controllers

// import (
// 	"net/http"
// 	"time"
// 	"uji-keterampilan-golang/config"
// 	"uji-keterampilan-golang/models"

// 	jwt "github.com/golang-jwt/jwt"
// 	"github.com/labstack/echo/v4"
// )

// type JWTClaims struct {
// 	Username string `json:"username"`
// 	jwt.StandardClaims
// }

// var jwtSecret = []byte("your-secret-key")

// func Register(c echo.Context) error {
// 	user := new(models.User)
// 	if err := c.Bind(user); err != nil {
// 		return err
// 	}

// 	// Simpan pengguna ke database
// 	result := config.DB.Create(user)
// 	if result.Error != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"error": "Failed to register user",
// 		})
// 	}

// 	return c.JSON(http.StatusCreated, map[string]interface{}{
// 		"message": "User registered successfully",
// 	})
// }

// func Login(c echo.Context) error {
// 	username := c.FormValue("username")
// 	password := c.FormValue("password")

// 	// Validasi username dan password
// 	if checkCredentials(username, password) {
// 		// Membuat token JWT
// 		token, err := createToken(username)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 				"error": "Failed to create JWT token",
// 			})
// 		}

// 		return c.JSON(http.StatusOK, map[string]interface{}{
// 			"token": token,
// 		})
// 	}

// 	return echo.ErrUnauthorized
// }

// func checkCredentials(username, password string) bool {
// 	var user models.User
// 	result := config.DB.Where("username = ?", username).First(&user)
// 	if result.Error != nil || user.Password != password {
// 		return false
// 	}
// 	return true
// }

// func createToken(username string) (string, error) {
// 	// Menentukan masa berlaku token (1 jam dalam contoh ini)
// 	expirationTime := time.Now().Add(1 * time.Hour)

// 	// Membuat payload JWT
// 	claims := &JWTClaims{
// 		Username: username,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expirationTime.Unix(),
// 		},
// 	}

// 	// Membuat token JWT dengan menggunakan kunci rahasia
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	signedToken, err := token.SignedString(jwtSecret)
// 	if err != nil {
// 		return "", err
// 	}

// 	return signedToken, nil
// }

// // Handler yang membutuhkan autentikasi
// func SecretHandler(c echo.Context) error {
// 	username := c.Get("username").(string)
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "Hello, " + username + "! This is a secret message.",
// 	})
// }
