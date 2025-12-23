package main

import (
	"net/http"

	"omnicampus/api/internal/config"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.Load()

	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	e.Logger.Fatal(e.Start(":" + cfg.APIPort))
}
