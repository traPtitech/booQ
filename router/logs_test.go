package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"github.com/traPtitech/booQ/model"
)

func TestPostLogs(t *testing.T) {
	t.Parallel()

	t.Run("post log", func(t *testing.T) {
		assert := assert.New(t)
		e := echoSetupWithAdminUser()

		item, _ := model.CreateItem(model.Item{Name: "testItemForPostLog"})
		user, _ := model.GetUserByName("traP")
		testBodyLog := model.Log{
			OwnerID: int(user.ID),
			Type:    1,
			Purpose: "ログのポストのテストのPurposeです",
			DueDate: time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC),
		}

		reqBody, _ := json.Marshal(testBodyLog)
		req := httptest.NewRequest(echo.POST, "/api/items/"+strconv.Itoa(int(item.ID))+"/logs", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusCreated, rec.Code)

		log := model.Log{}
		_ = json.NewDecoder(rec.Body).Decode(&log)

		assert.Equal(testBodyLog.OwnerID, log.OwnerID)
		assert.Equal(testBodyLog.Type, log.Type)
		assert.Equal(testBodyLog.Purpose, log.Purpose)
		assert.Equal(testBodyLog.DueDate, log.DueDate)
		assert.Equal(int(user.ID), log.UserID)
		assert.Equal(int(item.ID), log.ItemID)
	})

}
