package router

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/traPtitech/booQ/model"
)

// GetItems GET /items
func GetItems(c echo.Context) error {
	ownerName := c.QueryParam("user")
	if ownerName != "" {
		user, err := model.GetUserByName(ownerName)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if user.ID == 0 {
			return c.JSON(http.StatusBadRequest, errors.New("指定してるNameが不正です"))
		}
		res, err := model.SearchItemByOwner(ownerName)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, res)
	}
	searchString := c.QueryParam("search")
	if searchString != "" {
		res, err := model.SearchItems(searchString)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, res)
	}
	rentalName := c.QueryParam("rental")
	if rentalName != "" {
		user, err := model.GetUserByName(rentalName)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if user.ID == 0 {
			return c.JSON(http.StatusBadRequest, errors.New("指定してるNameが不正です"))
		}
		res, err := model.SearchItemByRental(user.ID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, res)
	}
	res, err := model.GetItems()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, res)
}

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
	item, err := model.GetItemByID(uint(itemID))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, item)
}

// PutItem PUT /items/:id
func PutItem(c echo.Context) error {
	ID := c.Param("id")
	user := c.Get("user").(model.User)
	body := map[string]interface{}{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	itemID, err := strconv.Atoi(ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	item, err := model.GetItemByID(uint(itemID))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	err = model.CheckOwnsOrAdmin(&user, &item)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	item = model.UpdateItem(&item, &body, user.Admin)

	return c.JSON(http.StatusOK, item)
}

// DeleteItem DELETE /items/:id
func DeleteItem(c echo.Context) error {
	ID := c.Param("id")
	user := c.Get("user").(model.User)
	itemID, err := strconv.Atoi(ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if !user.Admin {
		return c.NoContent(http.StatusForbidden)
	}
	item, err := model.GetItemByID(uint(itemID))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	item, err = model.DestroyItem(item)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, item)
}

// PostOwners POST /items/:id/owners
func PostOwners(c echo.Context) error {
	ID := c.Param("id")
	me := c.Get("user").(model.User)
	body := model.RequestPostOwnersBody{}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
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
	item, err := model.GetItemByID(uint(itemID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if body.UserID > 2 && item.Type > 0 {
		return c.NoContent(http.StatusForbidden)
	}
	if item.Type == 1 {
		user, _ = model.GetUserByName("traP")
	}
	if item.Type == 2 {
		user, _ = model.GetUserByName("sienka")
	}
	// item.Type=0⇒個人、1⇒trap(id:1)所有、2⇒支援課(id:2)
	if item.Type != 0 && !me.Admin {
		return c.NoContent(http.StatusForbidden)
	}
	owner := model.Owner{
		UserID:     user.ID,
		Rentalable: body.Rentalable,
		Count:      body.Count,
	}
	if body.Count < 0 {
		return c.NoContent(http.StatusBadRequest)
	}
	res, err := model.RegisterOwner(owner, item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, res)
}

// PutOwners PUT /items/:id/owners
func PutOwners(c echo.Context) error {
	ID := c.Param("id")
	me := c.Get("user").(model.User)
	body := model.RequestPostOwnersBody{}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
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
	item, err := model.GetItemByID(uint(itemID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if body.UserID > 2 && item.Type > 0 {
		return c.NoContent(http.StatusForbidden)
	}
	if item.Type == 1 {
		user, _ = model.GetUserByName("traP")
	}
	if item.Type == 2 {
		user, _ = model.GetUserByName("sienka")
	}
	// item.Type=0⇒個人、1⇒trap(id:1)所有、2⇒支援課(id:2)
	if item.Type != 0 && !me.Admin {
		return c.NoContent(http.StatusForbidden)
	}
	owner := model.Owner{
		UserID:     user.ID,
		Rentalable: body.Rentalable,
		Count:      body.Count,
	}
	res, err := model.AddOwner(owner, item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, res)
}

// PostLikes POST /items/:id/likes
func PostLikes(c echo.Context) error {
	ID := c.Param("id")
	user := c.Get("user").(model.User)
	user, err := model.GetUserByName(user.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	itemID, err := strconv.Atoi(ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	item, err := model.GetItemByID(uint(itemID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	_, err = model.CreateLike(item.ID, user.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.NoContent(http.StatusCreated)
}

// PostLikes POST /items/:id/likes
func DeleteLikes(c echo.Context) error {
	ID := c.Param("id")
	user := c.Get("user").(model.User)
	user, err := model.GetUserByName(user.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	itemID, err := strconv.Atoi(ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	item, err := model.GetItemByID(uint(itemID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	_, err = model.CancelLike(item.ID, user.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.NoContent(http.StatusCreated)
}
