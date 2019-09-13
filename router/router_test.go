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

	if !model.HasUsersTable() {
		err = model.Migrate()
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
				IconFileID:  "099eed74-3ab3-4655-ac37-bc7df1139b3d",
				Admin:       false,
			}
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
				IconFileID:  "099eed74-3ab3-4655-ac37-bc7df1139b3d",
				Admin:       true,
			}
			c.Set("user", adminUser)
			return c, nil
		},
	}
	SetupRouting(e, client)
	return e
}
