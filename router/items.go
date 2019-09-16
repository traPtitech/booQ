package router

import (
	"net/http"
	"strconv"

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

// GetItem GET /items/:id
func GetItem(c echo.Context) error {
	ID := c.Param("id")
	itemID, err := strconv.Atoi(ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	item, err := model.GetItemByID(itemID)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	return c.JSON(http.StatusOK, item)
}

// PostOwners POST /items/:id/owners
func PostOwners(c echo.Context) error {
	ID := c.Param("id")
	me := c.Get("user").(model.User)
	body := model.RequestPostOwnersBody{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	itemID, err := strconv.Atoi(ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	user, err := model.GetUserByID(body.UserID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = model.CheckTargetedOrAdmin(me, user)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	item, err := model.GetItemByID(itemID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// item.Type=0⇒個人、1⇒trap(id:1)所有、2⇒支援課(id:2)
	if item.Type == 1 && user.Name != "traP" {
		return c.NoContent(http.StatusForbidden)
	}
	if item.Type == 2 && user.Name != "sienka" {
		return c.NoContent(http.StatusForbidden)
	}
	var owner model.Owner
	owner.Owner = user
	owner.Rentalable = body.Rentalable
	res, err := model.RegisterOwner(owner, item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, res)
}
