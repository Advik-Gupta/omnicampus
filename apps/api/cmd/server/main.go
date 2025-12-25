package main

import (
	"log"

	"github.com/labstack/echo/v4"

	"omnicampus/api/db"
	"omnicampus/api/routes"
)

func main() {
	db.Init()
	defer db.Pool.Close()

	e := echo.New()

	routes.RegisterRoutes(e)

	log.Fatal(e.Start(":8080"))
}
