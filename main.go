package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	// gorm mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/traPtitech/booQ/model"
	"github.com/traPtitech/booQ/router"
	"github.com/traPtitech/booQ/storage"
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

	// Storage
	if os.Getenv("OS_CONTAINER") != "" {
		// Swiftオブジェクトストレージ
		err := storage.SetSwiftStorage(
			os.Getenv("OS_CONTAINER"),
			os.Getenv("OS_USERNAME"),
			os.Getenv("OS_PASSWORD"),
			os.Getenv("OS_TENANT_NAME"),
			os.Getenv("OS_TENANT_ID"),
			os.Getenv("OS_AUTH_URL"),
		)
		if err != nil {
			panic(err)
		}
	} else {
		// ローカルストレージ
		dir := os.Getenv("UPLOAD_DIR")
		if dir == "" {
			dir = "./uploads"
		}
		err := storage.SetLocalStorage(dir)
		if err != nil {
			panic(err)
		}
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routing
	router.SetupRouting(e, router.CreateUserProvider(os.Getenv("DEBUG_USER_NAME")))

	// Start server
	e.Logger.Fatal(e.Start(":3001"))
}
