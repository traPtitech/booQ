package router

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/traPtitech/booQ/model"
)

// GetUsersMe GET /users/me
func GetUsersMe(c echo.Context) error {
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

// PutUsers PUT /users
func PutUsers(c echo.Context) error {
	req := model.User{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	user := c.Get("user").(model.User)
	res, err := model.UpdateUser(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if user.Name != req.Name && !user.Admin {

		return c.NoContent(http.StatusForbidden)

	} else if !user.Admin {

		if req.Name != user.Name {
			return c.NoContent(http.StatusForbidden)
		} else if req.Admin {
			return c.NoContent(http.StatusForbidden)
		} else {
			return c.JSON(http.StatusOK, res)
		}

	} else {

		return c.JSON(http.StatusOK, res)

	}
}
