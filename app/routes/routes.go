package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/fazarth/inventory_go/app/controllers"
)

func InitRoutes(e *echo.Echo, userController *controllers.UserController, authController *controllers.AuthController) {
	e.GET("/users", userController.GetAllUsers)
	e.POST("/users", userController.CreateUser)
	e.GET("/users/:id", userController.GetUserByID)
	e.PUT("/users/:id", userController.UpdateUser)
	e.DELETE("/users/:id", userController.DeleteUser)

	e.POST("/login", authController.Login)
}
