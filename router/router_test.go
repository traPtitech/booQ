package router

import (
	"os"
	"testing"

	"github.com/labstack/echo"

	"github.com/traPtitech/booQ/model"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestMain(m *testing.M) {
	dbSetup()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func dbSetup() {
	_, err := model.EstablishConnection()
	if err != nil {
		panic(err)
	}

	err = model.Migrate()
	if err != nil {
		panic(err)
	}
	testUser, _ := model.GetUserByName("testUser")
	if testUser.Name == "" {
		user := model.User{
			Name:        "testUser",
			DisplayName: "テストユーザー",
			Admin:       false,
		}
		_, err = model.CreateUser(user)
		if err != nil {
			panic(err)
		}
	}
}

func echoSetupWithUser() *echo.Echo {
	e := echo.New()
	client := createMockUserProvider(model.User{
		Name:        "testUser",
		DisplayName: "テストユーザー",
		Admin:       false,
	})
	SetupRouting(e, client)
	SetValidator(e)
	return e
}

func echoSetupWithAdminUser() *echo.Echo {
	e := echo.New()
	client := createMockUserProvider(model.User{
		Name:        "traP",
		DisplayName: "traP",
		Admin:       true,
	})
	SetupRouting(e, client)
	SetValidator(e)
	return e
}
