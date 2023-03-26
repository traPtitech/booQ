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

func TestPostComments(t *testing.T) {
	testRequestBody := model.RequestPostCommentBody{
		Text: "testPostCommentsText",
	}
	testRequestInvalidBody := model.RequestPostCommentBody{
		Text: "",
	}
	item, _ := model.CreateItem(model.Item{Name: "testPostCommentsItem"})

	t.Run("admin user", func(t *testing.T) {
		assert := assert.New(t)
		e := echoSetupWithAdminUser()

		reqBody, _ := json.Marshal(testRequestBody)
		req := httptest.NewRequest(echo.POST, "/api/items/"+strconv.Itoa(int(item.ID))+"/comments", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusCreated, rec.Code)

		comment := model.Comment{}
		_ = json.NewDecoder(rec.Body).Decode(&comment)

		assert.Equal(testRequestBody.Text, comment.Text)
		assert.Equal(item.ID, comment.ItemID)
		assert.Equal("traP", comment.User.Name)
	})

	t.Run("admin user/validation error", func(t *testing.T) {
		assert := assert.New(t)
		e := echoSetupWithAdminUser()

		reqBody, _ := json.Marshal(testRequestInvalidBody)
		req := httptest.NewRequest(echo.POST, "/api/items/"+strconv.Itoa(int(item.ID))+"/comments", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusBadRequest, rec.Code)
	})
}
