package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/fazarth/inventory_go/controller"
	"github.com/fazarth/inventory_go/middleware"
)

func StartServer() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	router := NewRouter(e)
	router.SetupRoutes()

	// Start server
	e.Start(":8000")
}