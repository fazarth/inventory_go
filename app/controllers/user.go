package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/fazarth/inventory_go/app/helpers"
	"github.com/fazarth/inventory_go/app/models"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

func (c *UserController) GetAllUsers(ctx echo.Context) error {
	var users []models.User
	result := c.DB.Find(&users)
	if result.Error != nil {
		return helpers.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get users")
	}
	return helpers.SuccessResponse(ctx, http.StatusOK, users)
}

func (c *UserController) CreateUser(ctx echo.Context) error {
	user := new(models.User)
	if err := ctx.Bind(user); err != nil {
		return helpers.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload")
	}
	result := c.DB.Create(&user)
	if result.Error != nil {
		return helpers.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create user")
	}
	return helpers.SuccessResponse(ctx, http.StatusCreated, user)
}

func (c *UserController) GetUserByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return helpers.ErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID")
	}
	var user models.User
	result := c.DB.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return helpers.ErrorResponse(ctx, http.StatusNotFound, "User not found")
		}
		return helpers.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get user")
	}
	return helpers.SuccessResponse(ctx, http.StatusOK, user)
}

func (c *UserController) UpdateUser(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return helpers.ErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID")
	}
	var user models.User
	result := c.DB.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return helpers.ErrorResponse(ctx, http.StatusNotFound, "User not found")
		}
		return helpers.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get user")
	}
	if err := ctx.Bind(&user); err != nil {
		return helpers.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload")
	}
	result = c.DB.Save(&user)
	if result.Error != nil {
		return helpers.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to update user")
	}
	return helpers.SuccessResponse(ctx, http.StatusOK, user)
}

func (c *UserController) DeleteUser(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return helpers.ErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID")
	}
	var user models.User
	result := c.DB.Delete(&user, id)
	if result.Error != nil {
		return helpers.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete user")
	}
	return helpers.SuccessResponse(ctx, http.StatusOK, nil)
}
