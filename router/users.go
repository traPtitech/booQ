package router

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/traPtitech/booQ/model"
)

// GetUsersMe GET /users/me
func GetUsersMe(c echo.Context) error {
	user := c.Get("user").(model.User)
	res, err := model.GetUserByName(user.Name)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
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

// PutUsers PUT /users
func PutUsers(c echo.Context) error {
	req := model.RequestPutUsersBody{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user := c.Get("user").(model.User)
	prevUser, err := model.GetUserByName(req.Name)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	if !user.Admin {
		return c.NoContent(http.StatusForbidden)
	}
	if user.Admin == prevUser.Admin {
		return c.NoContent(http.StatusBadRequest)
	}

	prevUser.Admin = req.Admin
	res, err := model.UpdateUser(prevUser)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, res)
}
