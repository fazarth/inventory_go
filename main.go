package main

import (
	"github.com/fazarth/inventory_go/config"
)

func main() {
	// Initialize the database connection
	db := config.DBInstance()

	// Migrate the database
	db.AutoMigrate(
		&model.User{},
		&model.Product{},
		&model.Category{},
	)

	// Start the server
	config.StartServer()
}