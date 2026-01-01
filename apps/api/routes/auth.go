package routes

import (
	"omnicampus/api/internal/handlers"

	"github.com/labstack/echo/v4"
)

func authRoutes(api *echo.Group) {
	auth := api.Group("/auth")
	auth.POST("/request-otp", handlers.RequestOTP)
	auth.POST("/verify-otp", handlers.VerifyOTP)
	auth.POST("/set-password", handlers.SetPassword)
	auth.POST("/login", handlers.LoginUser)
	auth.GET("/me", handlers.MeHandler)
}