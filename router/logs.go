package router

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/traPtitech/booQ/model"
)

// PostLogs POST /items/:id/logs
func PostLogs(c echo.Context) error {
	ID := c.Param("id")
	user := c.Get("user").(model.User)
	user, _ = model.GetUserByName(user.Name)
	body := model.RequestPostLogsBody{}
	if err := c.Bind(&body); err != nil {
		return err
	}

	itemID, err := strconv.Atoi(ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	log := model.Log{
		ItemID:  itemID,
		UserID:  int(user.ID),
		OwnerID: body.OwnerID,
		Type:    body.Type,
		Purpose: body.Purpose,
		DueDate: body.DueDate,
	}

	res, err := model.CreateLog(log)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, res)
}
