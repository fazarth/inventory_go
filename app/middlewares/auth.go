package middlewares

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authHeader := ctx.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing authorization header")
		}

		tokenString := strings.Split(authHeader, " ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("your-secret-key"), nil
		})
		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to parse token claims")
		}

		username := claims["username"].(string)

		// Simulasi validasi user dari database
		if username != "admin" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		ctx.Set("username", username)

		return next(ctx)
	}
}

func GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username

	// Set expiration time
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Generate encoded token
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
