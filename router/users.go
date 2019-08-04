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

// PutUser PUT /users
func PutUser(c echo.Context) error {
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
	} else if res.Admin != req.Admin && !user.Admin {
		return c.NoContent(http.StatusForbidden)
	} else {
		return c.JSON(http.StatusOK, res)
	}

}
