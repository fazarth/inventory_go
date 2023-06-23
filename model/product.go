package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Price    float64
	Quantity int
	Category Category
}

type Category struct {
	gorm.Model
	Name string `gorm:"not null"`
}

type ProductResponse struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Price     float64    `json:"price"`
	Quantity  int        `json:"quantity"`
	Category  Category   `json:"category"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}