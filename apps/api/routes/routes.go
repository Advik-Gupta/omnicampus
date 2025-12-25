package routes

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	api := e.Group("")


	// Health check (optional)
	api.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok bro"})
	})

	devRoutes(api) // seeder routes
	authRoutes(api) // auth routes
}
