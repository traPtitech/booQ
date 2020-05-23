package router

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"

	"github.com/traPtitech/booQ/model"
)

var baseURL = "https://q.trap.jp/api/1.0"

// Traq traQに接続する用のclient
type Traq interface {
	GetUsersMe(c echo.Context) (echo.Context, error)
	MiddlewareAuthUser(next echo.HandlerFunc) echo.HandlerFunc
}

// TraqClient 本番用のclient
type TraqClient struct {
	Traq
}

// MockTraqClient テスト用のモックclient
type MockTraqClient struct {
	Traq
	MockGetUsersMe func(c echo.Context) (echo.Context, error)
}

// GetUsersMe 本番用のGetUsersMe
func (client *TraqClient) GetUsersMe(c echo.Context) (echo.Context, error) {
	res := c.Request().Header.Get("X-Showcase-User")
	if res == "" {
		return c, errors.New("認証に失敗しました(Headerに必要な情報が存在しません)")
	}
	user, _ := model.GetUserByName(res)
	if user.Name == "" {
		user, _ = model.CreateUser(model.User{Name: res})
	}
	c.Set("user", user)
	return c, nil
}

// GetUsersMe テスト用のGetUsersMe
func (client *MockTraqClient) GetUsersMe(c echo.Context) (echo.Context, error) {
	return client.MockGetUsersMe(c)
}

// MiddlewareAuthUser APIにアクセスしたユーザーの情報をセットする
func (client *TraqClient) MiddlewareAuthUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c, err := client.GetUsersMe(c)
		if err != nil {
			return c.String(http.StatusUnauthorized, err.Error())
		}
		return next(c)
	}
}

// MiddlewareAuthUser APIにアクセスしたユーザーの情報をセットする
func (client *MockTraqClient) MiddlewareAuthUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c, err := client.GetUsersMe(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}
		return next(c)
	}
}
