package main

import (
	"github.com/fazarth/inventory_go/app/controllers"
	"github.com/fazarth/inventory_go/app/middlewares"
	"github.com/fazarth/inventory_go/config"
	"github.com/fazarth/inventory_go/routes"
)

func main() {
	db := config.InitDB()
	defer db.Close()

	e := config.InitEcho()

	routes.InitRoutes(e, controllers.NewUserController(db), controllers.NewAuthController())

	e.Start(":8080")
}