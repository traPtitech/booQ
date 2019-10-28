package router

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/traPtitech/booQ/model"
)

// PostTags POST /tags
func PostTags(c echo.Context) error {
	body := model.RequestPostTagsBody{}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	tag, err := model.CreateTag(body.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusCreated, tag)
}

// PostTags2Item POST /items/:id/tags
func PostTags2Item(c echo.Context) error {
	ID := c.Param("id")
	itemID, err := strconv.Atoi(ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	body := model.RequestPostTags2ItemBody{}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	tags := []model.Tag{}
	for _, id := range body.ID {
		tag, err := model.AttachTag(id, uint(itemID))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		tags = append(tags, tag)
	}
	return c.JSON(http.StatusOK, tags)
}

// DeleteTag DELETE /items/:itemId/tags/:tagId
func DeleteTag(c echo.Context) error {
	itemId := c.Param("itemId")
	tagId := c.Param("tagId")
	itemID, err := strconv.Atoi(itemId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	tagID, err := strconv.Atoi(tagId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	tag, err := model.GetTagByID(uint(tagID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	_, err = model.RemoveTag(tag, uint(itemID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.NoContent(http.StatusOK)
}
