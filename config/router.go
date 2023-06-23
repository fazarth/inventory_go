package config

import (
	"github.com/labstack/echo/v4"
	"github.com/fazarth/inventory_go/controller"
	"github.com/fazarth/inventory_go/middleware"
)

type Router struct {
	Echo *echo.Echo
}

func NewRouter(e *echo.Echo) *Router {
	return &Router{
		Echo: e,
	}
}

func (r *Router) SetupRoutes() {
	r.Echo.POST("/register", controller.CreateUser)
	r.Echo.POST("/login", controller.Login)

	api := r.Echo.Group("/api")
	api.Use(middleware.JWTMiddleware)

	// Product routes
	api.POST("/products", controller.CreateProduct)
	api.GET("/products", controller.GetProducts)
}