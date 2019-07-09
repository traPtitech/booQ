package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	
	// gorm mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/traPtitech/booQ/model"
	"github.com/traPtitech/booQ/router"
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

	err = model.Migrate()
	if err != nil {
		panic(err)
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routing
	router.SetupRouting(e)

	// Start server
	e.Logger.Fatal(e.Start(":3001"))
}
