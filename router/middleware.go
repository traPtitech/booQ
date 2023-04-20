package router

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"

	"github.com/traPtitech/booQ/model"
)

// UserProvider traQに接続する用のclient
type UserProvider struct {
	AuthUser func(c echo.Context) (echo.Context, error)
}

// MiddlewareAuthUser APIにアクセスしたユーザーの情報をセットする
func (client *UserProvider) MiddlewareAuthUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c, err := client.AuthUser(c)
		if err != nil {
			return c.String(http.StatusUnauthorized, err.Error())
		}
		return next(c)
	}
}

// MiddlewareAdmin Admin以外を弾くmiddleware
func MiddlewareAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(model.User)
		if !user.Admin {
			return c.NoContent(http.StatusForbidden)
		}
		return next(c)
	}
}

// MiddlewareItemSocial ItemがPersonalItemでない場合はAdmin以外を弾くmiddleware
func MiddlewareItemSocial(getItem func(c echo.Context) model.Item) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			item := getItem(c)
			user := c.Get("user").(model.User)
			if item.Type != model.PersonalItem && !user.Admin {
				return c.NoContent(http.StatusForbidden)
			}
			return next(c)
		}
	}
}

func CreateUserProvider(debugUserName string) *UserProvider {
	return &UserProvider{AuthUser: func(c echo.Context) (echo.Context, error) {
		res := debugUserName
		if debugUserName == "" {
			res = c.Request().Header.Get("X-Showcase-User")
			if res == "" {
				return c, errors.New("認証に失敗しました(Headerに必要な情報が存在しません)")
			}
		}
		user, _ := model.GetUserByName(res)
		if user.Name == "" {
			user, _ = model.CreateUser(model.User{Name: res})
		}
		c.Set("user", user)
		return c, nil
	}}
}

func createMockUserProvider(user model.User) *UserProvider {
	return &UserProvider{AuthUser: func(c echo.Context) (echo.Context, error) {
		user, _ = model.GetUserByName(user.Name)
		c.Set("user", user)
		return c, nil
	}}
}
