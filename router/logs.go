package router

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"

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

	itemIDb, err := strconv.Atoi(ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	itemID := uint(itemIDb)
	item, _ := model.GetItemByID(itemID)
	var itemCount int
	var exist bool
	for _, owner := range item.Owners {
		if owner.UserID == body.OwnerID {
			if !owner.Rentalable {
				return c.NoContent(http.StatusForbidden)
			}
			itemCount = owner.Count
			exist = true
		}
	}
	if !exist {
		fmt.Print("指定のUserはそのItemを持っていません")
		return c.JSON(http.StatusBadRequest, errors.New("指定のUserはそのItemを持っていません"))
	}
	latestLog, err := model.GetLatestLog(item.Logs, body.OwnerID)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	dueDate, err := time.Parse("2006-01-02", body.DueDate)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	log := model.Log{
		ItemID:  itemID,
		UserID:  user.ID,
		OwnerID: body.OwnerID,
		Type:    body.Type,
		Purpose: body.Purpose,
		DueDate: dueDate,
	}
	var res model.Log
	if body.Type == 0 {
		if (latestLog.ItemID == 0) && (itemCount-body.Count < 0) {
			fmt.Print(itemCount - body.Count)
			fmt.Println("Rental超過1")
			return c.NoContent(http.StatusBadRequest)
		} else {
			if (latestLog.ItemID != 0) && latestLog.Count-body.Count < 0 {
				fmt.Println("Rental超過2")
				return c.NoContent(http.StatusBadRequest)
			}
		}
		if (latestLog == model.Log{}) {
			log.Count = itemCount - body.Count
		} else {
			log.Count = latestLog.Count - body.Count
		}
		rentalUser := model.RentalUser{
			UserID:  user.ID,
			OwnerID: body.OwnerID,
			Count:   body.Count * -1,
		}
		_, err = model.RentalItem(rentalUser, item)
		if err != nil {
			fmt.Print("rentalItemErr")
			return c.JSON(http.StatusBadRequest, err)
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
		log.Count = latestLog.Count + body.Count
		rentalUser := model.RentalUser{
			UserID:  user.ID,
			OwnerID: body.OwnerID,
			Count:   body.Count,
		}
		_, err = model.RentalItem(rentalUser, item)
		if err != nil {
			fmt.Print("rentalItemErr")
			return c.JSON(http.StatusBadRequest, err)
		}
	}
	res, err = model.CreateLog(log)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if res.ItemID == 0 {
		return c.NoContent(http.StatusBadRequest)
	}
	message := createMessage(log, item, user)
	_ = PostMessage(c, message)
	return c.JSON(http.StatusCreated, res)
}

func createMessage(log model.Log, item model.Item, user model.User) string {
	action := ""
	message := ""
	itemInfo := fmt.Sprintf("[%v](https://%v/items/%v)", item.Name, os.Getenv("HOST"), item.ID)
	if item.Type == 0 {
		purpose := ""
		if log.Type == 0 {
			action = "出"
			purpose = fmt.Sprintf("目的: %v", log.Purpose)
		} else {
			action = "入"
		}
		message = fmt.Sprintf("@%v \n%v\n%v × %v\n%v", user.Name, action, itemInfo, math.Abs(float64(log.Count)), purpose)
	} else {
		if log.Type == 0 {
			action = "借り"
		} else {
			action = "返し"
		}
		message = fmt.Sprintf("@%v が%vを%vました", user.Name, itemInfo, action)
	}
	return message
}
