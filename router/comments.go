package router

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"

	"github.com/traPtitech/booQ/model"
)

// PostComments POST /items/:id/comments
func PostComments(c echo.Context) error {
	ID := c.Param("id")
	itemID, err := strconv.Atoi(ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	item, err := model.GetItemByID(uint(itemID))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	user := c.Get("user").(model.User)
	user, err = model.GetUserByName(user.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	commentRequest := model.RequestPostCommentBody{}
	if err := c.Bind(&commentRequest); err != nil {
		return err
	}
	if err := c.Validate(&commentRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	comment := model.Comment{
		ItemID: uint(itemID),
		UserID: user.ID,
		User:   user,
		Text:   commentRequest.Text,
	}
	res, err := model.CreateComment(comment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	itemInfo := fmt.Sprintf("[%v](https://%v/items/%v)", item.Name, os.Getenv("HOST"), item.ID)
	message := fmt.Sprintf("### @%s がコメントを投稿しました\n%s\n%s", comment.User.Name, itemInfo, comment.Text)
	_ = PostMessage(c, message, false)
	return c.JSON(http.StatusCreated, res)
}

// GetComments GET /comments
func GetComments(c echo.Context) error {
	userName := c.QueryParam("user")
	if userName != "" {
		user, err := model.GetUserByName(userName)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		res, err := model.GetCommentsByUserID(user.ID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, res)
	}
	return c.JSON(http.StatusBadRequest, errors.New("require params `user`"))
}
