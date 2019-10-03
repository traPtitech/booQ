package router

import (
	"fmt"
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
		// return c.JSON(http.StatusBadRequest, err)
		return c.JSON(http.StatusBadRequest, model.Log{Purpose: "itemid"})
	}
	item, _ := model.GetItemByID(itemID)
	var itemCount int
	for _, owner := range item.Owners {
		if int(owner.OwnerID) == body.OwnerID {
			if !owner.Rentalable {
				return c.NoContent(http.StatusForbidden)
			}
			itemCount = owner.Count
		}
	}
	latestLog, err := model.GetLatestLog(itemID, body.OwnerID)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	var res model.Log
	if body.Type == 0 {
		if (latestLog.ItemID == 0) && (itemCount-body.Count < 0) {
			fmt.Println("Rental超過1")
			return c.NoContent(http.StatusBadRequest)
		} else {
			if (latestLog.ItemID != 0) && latestLog.Count-body.Count < 0 {
				fmt.Println("Rental超過2")
				return c.NoContent(http.StatusBadRequest)
			}
		}
		if (latestLog == model.Log{}) {
			log := model.Log{
				ItemID:  itemID,
				UserID:  int(user.ID),
				OwnerID: body.OwnerID,
				Type:    body.Type,
				Purpose: body.Purpose,
				DueDate: body.DueDate,
				Count:   itemCount - body.Count,
			}
			res, err = model.CreateLog(log)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
		} else {
			log := model.Log{
				ItemID:  itemID,
				UserID:  int(user.ID),
				OwnerID: body.OwnerID,
				Type:    body.Type,
				Purpose: body.Purpose,
				DueDate: body.DueDate,
				Count:   latestLog.Count - body.Count,
			}
			res, err = model.CreateLog(log)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
		}
	}
	if body.Type == 1 {
		if latestLog.ItemID == 0 {
			fmt.Println("Return超過1")
			return c.NoContent(http.StatusBadRequest)
		} else {
			if (latestLog.ItemID != 0) && itemCount-latestLog.Count-body.Count < 0 {
				fmt.Println("Return超過2")
				return c.NoContent(http.StatusBadRequest)
			}
		}
		log := model.Log{
			ItemID:  itemID,
			UserID:  int(user.ID),
			OwnerID: body.OwnerID,
			Type:    body.Type,
			Purpose: body.Purpose,
			DueDate: body.DueDate,
			Count:   latestLog.Count + body.Count,
		}
		res, err = model.CreateLog(log)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
	}
	if res.ItemID == 0 {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.JSON(http.StatusCreated, res)
}
