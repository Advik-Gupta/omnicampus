package routes

import (
	"os"

	"github.com/labstack/echo/v4"

	"omnicampus/api/internal/controllers"
)

func RegisterDevRoutes(e *echo.Echo) {
	if os.Getenv("ENV") != "dev" {
		return
	}

	dev := e.Group("/dev")

	dev.GET("/seed", controllers.SeedStudents)
}
