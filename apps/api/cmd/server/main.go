package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"omnicampus/api/internal/config"
	"omnicampus/api/pkg/utils"
	"omnicampus/api/routes"
)

func main() {
	cfg := config.Load()
	utils.InitDB()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://omnicampus-production.up.railway.app",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	routes.RegisterRoutes(e)

	log.Fatal(e.Start(":" + cfg.Port))
}
