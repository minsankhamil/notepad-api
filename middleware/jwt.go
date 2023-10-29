package middleware

import (
	"net/http"
	"notepad-api/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

var jwtToken = []byte("12345678")

func GenerateToken(user models.User) (string, error) {
	// Create the claims
	claims := jwt.MapClaims{
		"id_user": user.Id,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	// Create the token with the claims and the signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtToken)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func jwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, "Authorization token is missing")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtToken, nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, "Invalid or expired token")
		}

		// If the token is valid, you can access user data from the claims
		claims := token.Claims.(jwt.MapClaims)
		userID := claims["id_user"].(uint)

		// You can store the user ID in the context for use in your protected routes
		c.Set("userID", userID)

		return next(c)
	}
}
