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
	name := c.QueryParam("name")
	if name == "" {
		res := model.GetUsers()
		return c.JSON(http.StatusOK, res)
	}
	result, err := model.GetUserByName(name)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, result)
}
