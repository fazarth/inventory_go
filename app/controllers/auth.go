package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/fazarth/inventory_go/app/helpers"
	"github.com/fazarth/inventory_go/app/models"
	"github.com/fazarth/inventory_go/app/middlewares"
	"github.com/fazarth/inventory_go/config"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (c *AuthController) Login(ctx echo.Context) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	// Simulasi validasi user dari database
	if username != "admin" || password != "password" {
		return helpers.ErrorResponse(ctx, http.StatusUnauthorized, "Invalid username or password")
	}

	// Generate token JWT
	token, err := middlewares.GenerateJWT(username)
	if err != nil {
		return helpers.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to generate token")
	}

	// Simpan token ke database (opsional)
	// SimpanTokenKeDB(username, token)

	response := models.Token{
		Token: token,
	}

	return helpers.SuccessResponse(ctx, http.StatusOK, response)
}
