package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"github.com/traPtitech/booQ/model"

	"strconv"
)

func TestPostTags(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)
		e := echoSetupWithUser()

		testBody := model.RequestPostTagsBody{
			Name: "testPostTagsName",
		}
		reqBody, _ := json.Marshal(testBody)
		req := httptest.NewRequest(echo.POST, "/api/tags", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusCreated, rec.Code)

		tag := model.Tag{}
		_ = json.NewDecoder(rec.Body).Decode(&tag)

		assert.Equal(testBody.Name, tag.Name)
	})
}

func TestPostTags2Item(t *testing.T) {
	assert := assert.New(t)
	item, err := model.CreateItem(model.Item{Name: "testPostTags2ItemItem"})
	assert.NoError(err)
	tag1, err := model.CreateTag("testPostTags2ItemTag1")
	assert.NoError(err)
	tag2, err := model.CreateTag("testPostTags2ItemTag2")
	assert.NoError(err)
	tag3, err := model.CreateTag("testPostTags2ItemTag3")
	assert.NoError(err)
	testPostBody := model.RequestPostTags2ItemBody{}
	testPostBody.ID = append(testPostBody.ID, tag1.ID)
	testPostBody.ID = append(testPostBody.ID, tag2.ID)
	testPostBody.ID = append(testPostBody.ID, tag3.ID)

	t.Run("success", func(t *testing.T) {
		e := echoSetupWithAdminUser()

		reqBody, _ := json.Marshal(testPostBody)
		req := httptest.NewRequest(echo.POST, "/api/items/"+strconv.Itoa(int(item.ID))+"/tags", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusOK, rec.Code)

		tags := []model.Tag{}
		_ = json.NewDecoder(rec.Body).Decode(&tags)

		assert.Equal(tag1.Name, tags[0].Name)
		assert.Equal(tag2.Name, tags[1].Name)
		assert.Equal(tag3.Name, tags[2].Name)
	})
}

func TestDeleteTag(t *testing.T) {
	assert := assert.New(t)
	t.Run("fail", func(t *testing.T) {
		e := echoSetupWithAdminUser()
		item, err := model.CreateItem(model.Item{Name: "testDeleteTagFailItem"})
		assert.NoError(err)
		tag, err := model.CreateTag("testDeleteTagFailTag")
		assert.NoError(err)

		req := httptest.NewRequest(echo.DELETE, "/api/items/"+strconv.Itoa(int(item.ID))+"/tags/"+strconv.Itoa(int(tag.ID)), nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusBadRequest, rec.Code)
	})

	t.Run("success", func(t *testing.T) {
		e := echoSetupWithAdminUser()
		item, err := model.CreateItem(model.Item{Name: "testDeleteTagSuccessItem"})
		assert.NoError(err)
		tag, err := model.CreateTag("testDeleteTagSuccessTag")
		assert.NoError(err)
		_, err = model.AttachTag(tag.ID, item.ID)
		assert.NoError(err)

		req := httptest.NewRequest(echo.DELETE, "/api/items/"+strconv.Itoa(int(item.ID))+"/tags/"+strconv.Itoa(int(tag.ID)), nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusOK, rec.Code)
	})
}
