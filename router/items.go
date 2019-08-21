package router

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/traPtitech/booQ/model"
)

// PostItems POST /user/me
func PostItems(c echo.Context) error {
	item := model.Item{}
	if err := c.Bind(&item); err != nil {
		return err
	}
	res, err := model.CreateItem(item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, res)
}
