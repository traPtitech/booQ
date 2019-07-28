package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"

	"github.com/traPtitech/booQ/model"
)

var baseURL = "https://q.trap.jp/api/1.0"

// Traq traQに接続する用のclient
type Traq interface {
	GetUsersMe(c echo.Context) (echo.Context, error)
}

// TraqClient 本番用のclient
type TraqClient struct {
	client Traq
}

// MockTraqClient テスト用のモックclient
type MockTraqClient struct {
	Traq
	MockGetUsersMe func(c echo.Context) (echo.Context, error)
}

// GetUsersMe 本番用のGetUsersMe
func (client *TraqClient) GetUsersMe(c echo.Context) (echo.Context, error) {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c, c.NoContent(http.StatusUnauthorized)
	}
	req, _ := http.NewRequest("GET", baseURL+"/users/me", nil)
	req.Header.Set("Authorization", token)
	httpClient := new(http.Client)
	res, _ := httpClient.Do(req)
	if res.StatusCode != 200 {
		return c, c.NoContent(http.StatusUnauthorized)
	}
	body, _ := ioutil.ReadAll(res.Body)
	user := model.User{}
	err := json.Unmarshal(body, &user)
	if err != nil {
		return c, c.NoContent(http.StatusInternalServerError)
	}
	c.Set("user", user)
	return c, nil
}

// GetUsersMe テスト用のGetUsersMe
func (client *MockTraqClient) GetUsersMe(c echo.Context) (echo.Context, error) {
	return client.MockGetUsersMe(c)
}
