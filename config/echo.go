package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitEcho() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Tambahkan middleware autentikasi di sini
	e.Use(middlewares.AuthMiddleware)

	return e
}
