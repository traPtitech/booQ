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
	res := model.GetUsers()
	return c.JSON(http.StatusOK, res)
}

// PutUsers PUT /users
func PutUsers(c echo.Context) error {
	req := model.User{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user := c.Get("user").(model.User)
	err = model.CheckTargetedOrAdmin(user, req)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	if user.Name != req.Name && !user.Admin {
		return c.NoContent(http.StatusForbidden)
	}
	if !user.Admin && req.Admin {
		return c.NoContent(http.StatusForbidden)
	}
	res, err := model.UpdateUser(req)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, res)
}
