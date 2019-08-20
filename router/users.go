package router

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/traPtitech/booQ/model"
)

// GetUserMe GET /user/me
func GetUserMe(c echo.Context) error {
	user := c.Get("user").(model.User)
	res, err := model.GetUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if res.Name == "" {
		res, _ = model.CreateUser(user)
	}

	return c.JSON(http.StatusOK, res)
}

//GetUsers GET /users
func GetUsers(c echo.Context) error {
	req := model.RequestUserName{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	var res []model.User
	for _, value := range req.Name {
		result, err := model.GetUserByName(value)
		if err != nil {
			break
		}
		if result.Name == "" {
			continue
		}
		res = append(res, result)
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	} else {
		return c.JSON(http.StatusOK, res)
	}
}
