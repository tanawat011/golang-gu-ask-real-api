package users

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type loginModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var tokenSecret = "devx"

// UserLogin ...
func UserLogin(c echo.Context) error {
	data := new(loginModel)
	if err := c.Bind(data); err != nil {
		return err
	}

	// Throws unauthorized error
	if data.Username != "tanawat" || data.Password != "198777" {
		return echo.ErrUnauthorized
	}

	accessToken, err := getAccessToken()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"access_token": accessToken,
	})
}

func getAccessToken() (string, error) {

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Tanawat Pinthongpan"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	accessToken, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
