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

	item, _ := model.CreateItem(model.Item{Name: "testItemForPostLog"})
	itemRentalDenied, _ := model.CreateItem(model.Item{Name: "testItemForPostLogRentalDenied"})
	trap, _ := model.GetUserByName("traP")
	owner := model.Owner{
		OwnerID:    int(trap.ID),
		Rentalable: true,
		Count:      1,
	}
	ownerRentalDenied := model.Owner{
		OwnerID:    int(trap.ID),
		Rentalable: false,
		Count:      1,
	}
	_, _ = model.RegisterOwner(owner, item)
	_, _ = model.RegisterOwner(ownerRentalDenied, itemRentalDenied)
	testBodyLogRental1 := model.Log{
		OwnerID: int(trap.ID),
		Type:    0,
		Purpose: "ログのポストのテストのPurposeですrental1",
		DueDate: time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC),
		Count:   1,
	}
	testBodyLogRental2 := model.Log{
		OwnerID: int(trap.ID),
		Type:    0,
		Purpose: "ログのポストのテストのPurposeですrental2",
		DueDate: time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC),
		Count:   2,
	}
	testBodyLogReturn1 := model.Log{
		OwnerID: int(trap.ID),
		Type:    1,
		Purpose: "ログのポストのテストのPurposeですreturn1",
		DueDate: time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC),
		Count:   1,
	}
	testBodyLogReturn1RentalDenied := model.Log{
		OwnerID: int(trap.ID),
		Type:    1,
		Purpose: "ログのポストのテストのPurposeですreturn1rentalDenied",
		DueDate: time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC),
		Count:   1,
	}

	t.Run("failed", func(t *testing.T) {
		assert := assert.New(t)
		e := echoSetupWithAdminUser()

		reqBody, _ := json.Marshal(testBodyLogRental2)
		req := httptest.NewRequest(echo.POST, "/api/items/"+strconv.Itoa(int(item.ID))+"/logs", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusBadRequest, rec.Code)

		reqBody, _ = json.Marshal(testBodyLogReturn1)
		req = httptest.NewRequest(echo.POST, "/api/items/"+strconv.Itoa(int(item.ID))+"/logs", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec = httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusBadRequest, rec.Code)

		reqBody, _ = json.Marshal(testBodyLogReturn1RentalDenied)
		req = httptest.NewRequest(echo.POST, "/api/items/"+strconv.Itoa(int(itemRentalDenied.ID))+"/logs", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec = httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusForbidden, rec.Code)
	})

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)
		e := echoSetupWithAdminUser()

		reqBody, _ := json.Marshal(testBodyLogRental1)
		req := httptest.NewRequest(echo.POST, "/api/items/"+strconv.Itoa(int(item.ID))+"/logs", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusCreated, rec.Code)

		log := model.Log{}
		_ = json.NewDecoder(rec.Body).Decode(&log)

		assert.Equal(testBodyLogRental1.OwnerID, log.OwnerID)
		assert.Equal(testBodyLogRental1.Type, log.Type)
		assert.Equal(testBodyLogRental1.Purpose, log.Purpose)
		assert.Equal(testBodyLogRental1.DueDate, log.DueDate)
		assert.Equal(int(trap.ID), log.UserID)
		assert.Equal(int(item.ID), log.ItemID)

		reqBody, _ = json.Marshal(testBodyLogReturn1)
		req = httptest.NewRequest(echo.POST, "/api/items/"+strconv.Itoa(int(item.ID))+"/logs", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec = httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusCreated, rec.Code)

		log = model.Log{}
		_ = json.NewDecoder(rec.Body).Decode(&log)

		assert.Equal(testBodyLogReturn1.OwnerID, log.OwnerID)
		assert.Equal(testBodyLogReturn1.Type, log.Type)
		assert.Equal(testBodyLogReturn1.Purpose, log.Purpose)
		assert.Equal(testBodyLogReturn1.DueDate, log.DueDate)
		assert.Equal(int(trap.ID), log.UserID)
		assert.Equal(int(item.ID), log.ItemID)
	})

}
