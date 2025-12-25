package controllers

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"omnicampus/api/db"
)

func SeedStudents(c echo.Context) error {
	students := []struct {
		Name  string
		RegNo string
		DOB   string
		Email string
		Phone string
	}{
		{"User One", "21BCE0001", "2003-01-01", "user1@vit.edu", "9999999991"},
		{"User Two", "21BCE0002", "2003-02-02", "user2@vit.edu", "9999999992"},
		{"User Three", "21BCE0003", "2003-03-03", "user3@vit.edu", "9999999993"},
		{"User Four", "21BCE0004", "2003-04-04", "user4@vit.edu", "9999999994"},
		{"User Five", "21BCE0005", "2003-05-05", "user5@vit.edu", "9999999995"},
	}

	ctx := c.Request().Context()

	for _, s := range students {
		dob, err := time.Parse("2006-01-02", s.DOB)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid DOB format",
			})
		}

		_, err = db.Pool.Exec(
			ctx,
			`
			INSERT INTO student (
				id,
				name,
				register_number,
				dob,
				email,
				password,
				phone,
				timetable_id,
				courses_ids
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7, NULL, $8)
			ON CONFLICT (email) DO NOTHING
			`,
			uuid.New(),        // Google UUID
			s.Name,
			s.RegNo,
			dob,
			s.Email,
			"hashed_password", // dummy password
			s.Phone,
			[]uuid.UUID{},     // empty UUID array
		)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]string{
		"status": "dummy students seeded successfully",
	})
}
