package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"omnicampus/api/pkg/utils"
	"omnicampus/api/routes"
)

func main() {
	utils.LoadConfig()
	utils.InitDB()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	routes.RegisterRoutes(e)

	log.Fatal(e.Start(":8080"))
}
