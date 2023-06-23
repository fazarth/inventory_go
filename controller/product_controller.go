package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/fazarth/inventory_go/config"
	"github.com/fazarth/inventory_go/model"
	"gorm.io/gorm"
)

// CreateProduct creates a new product
func CreateProduct(c echo.Context) error {
	productRequest := new(model.Product)
	if err := c.Bind(productRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}
	if err := c.Validate(productRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	db := config.DBInstance()

	if err := db.Create(&productRequest).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create product"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Product created successfully"})
}

// GetProducts returns all products
func GetProducts(c echo.Context) error {
	db := config.DBInstance()

	var products []model.Product
	if err := db.Preload("Category").Find(&products).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to query database"})
	}

	response := make([]model.ProductResponse, len(products))
	for i, product := range products {
		response[i] = model.ProductResponse{
			ID:        product.ID,
			Name:      product.Name,
			Price:     product.Price,
			Quantity:  product.Quantity,
			Category:  product.Category,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		}
	}

	return c.JSON(http.StatusOK, response)
}