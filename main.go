package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/traPtitech/booQ/model"
)

func main() {

	db, err := model.EstablishConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if os.Getenv("BOOQ_ENV") == "development" {
		db.LogMode(true)
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/api/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	// Start server
	e.Logger.Fatal(e.Start(":3001"))
}
