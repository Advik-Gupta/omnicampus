package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"omnicampus/api/internal/controllers"
)

func RequestOTP(c echo.Context) error {
	var req struct {
		Email string `json:"email"`
	}

	if err := c.Bind(&req); err != nil || req.Email == "" {
		return c.JSON(http.StatusBadRequest, "invalid email")
	}

	err := controllers.SendOTP(req.Email)
	if err != nil {
		switch err.Error() {
		case "user not found":
			return c.JSON(http.StatusNotFound, "user does not exist")
		case "otp already sent":
			return c.JSON(http.StatusOK, "OTP already sent. Please check your email.")
		default:
			return c.JSON(http.StatusInternalServerError, "failed to send otp")
		}
	}

	return c.JSON(http.StatusOK, "OTP sent")
}

func VerifyOTP(c echo.Context) error {
	var req struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	token, err := controllers.VerifyOTP(req.Email, req.OTP)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "invalid or expired otp")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}