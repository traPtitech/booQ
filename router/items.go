package router

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/traPtitech/booQ/model"
)

// PostItems POST /items
func PostItems(c echo.Context) error {
	user := c.Get("user").(model.User)
	item := model.Item{}
	if err := c.Bind(&item); err != nil {
		return err
	}
	// item.Type=0⇒個人、1⇒trap所有、2⇒支援課
	if item.Type != 0 && !user.Admin {
		return c.NoContent(http.StatusForbidden)
	}
	res, err := model.CreateItem(item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, res)
}
