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
	item, _ := model.CreateItem(model.Item{Name: "testPostTags2ItemItem"})
	tag1, _ := model.CreateTag("testPostTags2ItemTag1")
	tag2, _ := model.CreateTag("testPostTags2ItemTag2")
	tag3, _ := model.CreateTag("testPostTags2ItemTag3")
	testPostBody := model.RequestPostTags2ItemBody{}
	testPostBody.ID = append(testPostBody.ID, tag1.ID)
	testPostBody.ID = append(testPostBody.ID, tag2.ID)
	testPostBody.ID = append(testPostBody.ID, tag3.ID)

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)
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
	t.Run("fail", func(t *testing.T) {
		assert := assert.New(t)
		e := echoSetupWithAdminUser()
		item, _ := model.CreateItem(model.Item{Name: "testDeleteTagFailItem"})
		tag, _ := model.CreateTag("testDeleteTagFailTag")

		req := httptest.NewRequest(echo.DELETE, "/api/items/"+strconv.Itoa(int(item.ID))+"/tags/"+strconv.Itoa(int(tag.ID)), nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusBadRequest, rec.Code)
	})

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)
		e := echoSetupWithAdminUser()
		item, _ := model.CreateItem(model.Item{Name: "testDeleteTagSuccessItem"})
		tag, _ := model.CreateTag("testDeleteTagSuccessTag")
		_, err := model.AttachTag(tag.ID, item.ID)
		assert.NoError(err)

		req := httptest.NewRequest(echo.DELETE, "/api/items/"+strconv.Itoa(int(item.ID))+"/tags/"+strconv.Itoa(int(tag.ID)), nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusOK, rec.Code)
	})
}
