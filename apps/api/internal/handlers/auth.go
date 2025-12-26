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
		case "user already onboarded":
			return c.JSON(http.StatusBadRequest, "user already onboarded")
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

	if err := controllers.VerifyOTP(req.Email, req.OTP); err != nil {
		switch err.Error() {
		case "invalid or expired otp":
			return c.JSON(http.StatusUnauthorized, "invalid or expired otp")
		case "user already onboarded":
			return c.JSON(http.StatusBadRequest, "user already onboarded")
		default:
			return c.JSON(http.StatusInternalServerError, "failed to verify otp")
		}
	}

	return c.JSON(http.StatusOK, map[string]bool{
		"verified": true,
	})
}

func SetPassword(c echo.Context) error {
	var req struct {
		Email       string `json:"email"`
		NewPassword string `json:"password"`
	}

	if err := c.Bind(&req); err != nil || req.Email == "" || req.NewPassword == "" {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	err := controllers.SetPassword(req.Email, req.NewPassword)
	if err != nil {
		switch err.Error() {
		case "invalid or expired token":
			return c.JSON(http.StatusUnauthorized, "invalid or expired token")
		case "user already onboarded":
			return c.JSON(http.StatusBadRequest, "user already onboarded")
		default:
			return c.JSON(http.StatusInternalServerError, "failed to set password")
		}
	}

	return c.JSON(http.StatusOK, "password set successfully")
}

func LoginUser (c echo.Context) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&req); err != nil || req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	token, err := controllers.LoginUser(req.Email, req.Password)
	if err != nil {
		switch err.Error() {
		case "user not found":
			return c.JSON(http.StatusNotFound, "user does not exist")
		case "incorrect password":
			return c.JSON(http.StatusUnauthorized, "incorrect password")
		default:
			return c.JSON(http.StatusInternalServerError, "failed to login")
		}
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}