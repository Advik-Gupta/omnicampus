package routes

import (
	"os"

	"github.com/labstack/echo/v4"

	"omnicampus/api/internal/controllers"
)

func devRoutes(api *echo.Group) {
	if os.Getenv("ENV") != "dev" {
		return
	}

	dev := api.Group("/dev")

	dev.GET("/seed", controllers.SeedStudents)
}
