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
	client := &MockTraqClient{
		MockGetUsersMe: func(c echo.Context) (echo.Context, error) {
			user := model.User{
				Name:        "testUser",
				DisplayName: "テストユーザー",
				Admin:       false,
			}
			user, _ = model.GetUserByName(user.Name)
			c.Set("user", user)
			return c, nil
		},
	}
	SetupRouting(e, client)
	return e
}

func echoSetupWithAdminUser() *echo.Echo {
	e := echo.New()
	client := &MockTraqClient{
		MockGetUsersMe: func(c echo.Context) (echo.Context, error) {
			adminUser := model.User{
				Name:        "traP",
				DisplayName: "traP",
				Admin:       true,
			}
			adminUser, _ = model.GetUserByName(adminUser.Name)
			c.Set("user", adminUser)
			return c, nil
		},
	}
	SetupRouting(e, client)
	return e
}
