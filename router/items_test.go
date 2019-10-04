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

func TestPostItems(t *testing.T) {
	t.Parallel()

	testBodyTrap := model.Item{
		Name:        "testTrapItem",
		Type:        1,
		Code:        "1920093013000",
		Description: "これは備品のテストです",
		ImgURL:      "http://example.com/testTrap.jpg",
	}

	testBodyKojin := model.Item{
		Name:        "testKojinItem",
		Type:        0,
		Code:        "9784049123944",
		Description: "これは個人所有物のテストです",
		ImgURL:      "http://example.com/testKojin.jpg",
	}

	t.Run("admin user", func(t *testing.T) {
		assert := assert.New(t)
		e := echoSetupWithAdminUser()

		reqBody, _ := json.Marshal(testBodyTrap)
		req := httptest.NewRequest(echo.POST, "/api/items", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusCreated, rec.Code)

		item := model.Item{}
		_ = json.NewDecoder(rec.Body).Decode(&item)

		assert.Equal(testBodyTrap.Name, item.Name)
		assert.Equal(testBodyTrap.Type, item.Type)
		assert.Equal(testBodyTrap.Code, item.Code)
		assert.Equal(testBodyTrap.Description, item.Description)
		assert.Equal(testBodyTrap.ImgURL, item.ImgURL)
	})

	t.Run("not admin user", func(t *testing.T) {
		assert := assert.New(t)
		e := echoSetupWithUser()

		reqBody, _ := json.Marshal(testBodyTrap)
		req := httptest.NewRequest(echo.POST, "/api/items", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusForbidden, rec.Code)

		reqBody, _ = json.Marshal(testBodyKojin)
		req = httptest.NewRequest(echo.POST, "/api/items", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec = httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusCreated, rec.Code)

		item := model.Item{}
		_ = json.NewDecoder(rec.Body).Decode(&item)

		assert.Equal(testBodyKojin.Name, item.Name)
		assert.Equal(testBodyKojin.Type, item.Type)
		assert.Equal(testBodyKojin.Code, item.Code)
		assert.Equal(testBodyKojin.Description, item.Description)
		assert.Equal(testBodyKojin.ImgURL, item.ImgURL)
	})
}

func TestPostOwners(t *testing.T) {
	t.Parallel()

	testBodyTrap := model.Item{
		Name:        "testTrapItem",
		Type:        1,
		Code:        "1920093013000",
		Description: "これは備品のテストです",
		ImgURL:      "http://example.com/testTrap.jpg",
	}

	testBodyKojin := model.Item{
		Name:        "testKojinItem",
		Type:        0,
		Code:        "9784049123944",
		Description: "これは個人所有物のテストです",
		ImgURL:      "http://example.com/testKojin.jpg",
	}

	trap, _ := model.GetUserByName("traP")

	testOwnerTrap := model.RequestPostOwnersBody{
		UserID:     int(trap.ID),
		Rentalable: true,
		Count:      1,
	}

	t.Run("admin user", func(t *testing.T) {
		assert := assert.New(t)
		e := echoSetupWithAdminUser()

		reqBody, _ := json.Marshal(testBodyTrap)
		req := httptest.NewRequest(echo.POST, "/api/items", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusCreated, rec.Code)

		item := model.Item{}
		_ = json.NewDecoder(rec.Body).Decode(&item)

		createdBihin, _ := model.GetItemByName(item.Name)
		bihinID := int(createdBihin.ID)
		reqBody, _ = json.Marshal(testOwnerTrap)
		req = httptest.NewRequest(echo.POST, "/api/items/"+strconv.Itoa(bihinID)+"/owners", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec = httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusOK, rec.Code)

		item = model.Item{}
		_ = json.NewDecoder(rec.Body).Decode(&item)

		assert.Equal(testBodyTrap.Name, item.Name)
		assert.Equal(trap.ID, item.Owners[0].OwnerID)
	})

	t.Run("not admin user", func(t *testing.T) {
		user, _ := model.GetUserByName("testUser")
		userID := int(user.ID)
		testOwnerKojin := model.RequestPostOwnersBody{
			UserID:     userID,
			Rentalable: true,
			Count:      1,
		}
		assert := assert.New(t)
		e := echoSetupWithUser()

		createdBihin, _ := model.GetItemByName("testTrapItem")
		bihinID := int(createdBihin.ID)
		reqBody, _ := json.Marshal(testOwnerKojin)
		req := httptest.NewRequest(echo.POST, "/api/items/"+strconv.Itoa(bihinID)+"/owners", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusForbidden, rec.Code)

		reqBody, _ = json.Marshal(testBodyKojin)
		req = httptest.NewRequest(echo.POST, "/api/items", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec = httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusCreated, rec.Code)

		item := model.Item{}
		_ = json.NewDecoder(rec.Body).Decode(&item)

		createdItem, _ := model.GetItemByName(item.Name)

		itemID := int(createdItem.ID)
		paramID := strconv.Itoa(itemID)
		targetAPI := "/api/items/" + paramID + "/owners"

		reqBody, _ = json.Marshal(testOwnerKojin)
		req = httptest.NewRequest(echo.POST, targetAPI, bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec = httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusOK, rec.Code)

		_ = json.NewDecoder(rec.Body).Decode(&item)

		assert.Equal(testBodyKojin.Name, item.Name)
		assert.Equal(user.ID, item.Owners[0].OwnerID)
	})
}
