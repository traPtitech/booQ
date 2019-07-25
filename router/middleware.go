package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"

	"github.com/traPtitech/booQ/model"
)

func middlewareAuthUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.NoContent(http.StatusForbidden)
		}
		req, _ := http.NewRequest("GET", "https://q.trap.jp/api/1.0/users/me", nil)
		req.Header.Set("Authorization", token)
		client := new(http.Client)
		res, _ := client.Do(req)
		if res.StatusCode != 200 {
			return c.NoContent(http.StatusForbidden)
		}
		body, _ := ioutil.ReadAll(res.Body)
		user := model.User{}
		err := json.Unmarshal(body, &user)
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		c.Set("user", user)
		return next(c)
	}
}
