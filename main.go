package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	appMiddleware "golang-gu-ask-real-api/middleware"
)

type loginModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	e := echo.New()

	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BasicAuth(appMiddleware.BasicAuth))

	// Login route
	e.POST("/login", login)

	e.Logger.Fatal(e.Start(":8080"))
}

// login ...
func login(c echo.Context) error {
	data := new(loginModel)
	if err := c.Bind(data); err != nil {
		return err
	}

	// Throws unauthorized error
	if data.Username != "tanawat" || data.Password != "198777" {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Tanawat Pinthongpan"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
