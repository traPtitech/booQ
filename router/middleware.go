package router

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

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

// MiddlewareBodyItemSocial リクエストボディから取得したItemがPersonalItemでない場合はAdmin以外を弾くmiddleware
func MiddlewareBodyItemSocial(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return next(c)
		}
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))
		item := model.Item{}
		if err = json.Unmarshal(body, &item); err != nil {
			return next(c)
		}
		user := c.Get("user").(model.User)
		if item.Type != model.PersonalItem && !user.Admin {
			return c.NoContent(http.StatusForbidden)
		}
		return next(c)
	}
}

// MiddlewareParamItemSocial パラメータから取得したItemがPersonalItemでない場合はAdmin以外を弾くmiddleware
func MiddlewareParamItemSocial(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		itemID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return next(c)
		}
		item, _ := model.GetItemByID(uint(itemID))
		user := c.Get("user").(model.User)
		if item.Type != model.PersonalItem && !user.Admin {
			return c.NoContent(http.StatusForbidden)
		}
		return next(c)
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
