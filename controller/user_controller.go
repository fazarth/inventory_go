package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/username/inventory-management/config"
	"github.com/username/inventory-management/model"
	"gorm.io/gorm"
)

// CreateUser creates a new user
func CreateUser(c echo.Context) error {
	userRequest := new(model.User)
	if err := c.Bind(userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}
	if err := c.Validate(userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	userRequest.HashPassword()

	db := config.DBInstance()

	if err := db.Create(&userRequest).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create user"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User created successfully"})
}

// Login authenticates a user and generates a JWT token
func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	db := config.DBInstance()
	var user model.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid username or password"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to query database"})
	}

	if err := model.VerifyPassword(user.Password, password); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid username or password"})
	}

	token, err := config.GenerateJWTToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to generate JWT token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}