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
		Name:        "testPostTrapItem",
		Type:        1,
		Code:        "1920093013000",
		Description: "これは備品のテストです",
		ImgURL:      "http://example.com/testTrap.jpg",
	}

	testBodyKojin := model.Item{
		Name:        "testPostKojinItem",
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

func TestDeleteItem(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	item, err := model.CreateItem(model.Item{Name: "testDeleteKojinItem"})
	assert.NoError(err)
	assert.NotEmpty(item)

	t.Run("admin user", func(t *testing.T) {
		e := echoSetupWithAdminUser()

		req := httptest.NewRequest(echo.DELETE, "/api/items/"+strconv.Itoa(int(item.ID)), nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusOK, rec.Code)

		nowItem, err := model.GetItemByID(item.ID)
		assert.Empty(nowItem)
		assert.Error(err)
	})

	t.Run("not admin user", func(t *testing.T) {
		e := echoSetupWithUser()

		req := httptest.NewRequest(echo.DELETE, "/api/items/"+strconv.Itoa(int(item.ID)), nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusForbidden, rec.Code)

		nowItem, err := model.GetItemByID(item.ID)
		assert.Empty(nowItem)
		assert.Error(err)
	})
}

func TestGetItem(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	t.Run("fail", func(t *testing.T) {
		e := echoSetupWithUser()

		req := httptest.NewRequest(echo.GET, "/api/items/999", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusNotFound, rec.Code)

		req = httptest.NewRequest(echo.GET, "/api/items/testfail", nil)
		req.Header.Set("Content-Type", "application/json")
		rec = httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusBadRequest, rec.Code)
	})

	t.Run("success", func(t *testing.T) {
		e := echoSetupWithUser()

		item, err := model.CreateItem(model.Item{Name: "testGetItemRouter"})
		assert.NoError(err)
		assert.NotEmpty(item)
		req := httptest.NewRequest(echo.GET, "/api/items/"+strconv.Itoa(int(item.ID)), nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusOK, rec.Code)

		gotItem := model.Item{}
		_ = json.NewDecoder(rec.Body).Decode(&gotItem)
		assert.Equal(item.Name, gotItem.Name)
	})
}

func TestGetItems(t *testing.T) {
	assert := assert.New(t)
	ownerUser, err := model.CreateUser(model.User{Name: "testGetItemsUser"})
	assert.NoError(err)

	t.Run("failed", func(t *testing.T) {
		e := echoSetupWithUser()

		req := httptest.NewRequest(echo.GET, "/api/items?rental=testUser", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusOK, rec.Code)
		gotItemsFail1 := []model.Item{}
		_ = json.NewDecoder(rec.Body).Decode(&gotItemsFail1)
		assert.Equal([]model.Item{}, gotItemsFail1)

		req = httptest.NewRequest(echo.GET, "/api/items?user=testGetItemsUser", nil)
		req.Header.Set("Content-Type", "application/json")
		rec = httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusOK, rec.Code)
		gotItemsFail2 := []model.Item{}
		_ = json.NewDecoder(rec.Body).Decode(&gotItemsFail2)
		assert.Equal([]model.Item{}, gotItemsFail2)
	})

	t.Run("success", func(t *testing.T) {
		e := echoSetupWithUser()

		item, err := model.CreateItem(model.Item{Name: "testGetItemsItem"})
		assert.NoError(err)
		owner := model.Owner{
			UserID:     ownerUser.ID,
			Rentalable: true,
			Count:      1,
		}
		_, err = model.RegisterOwner(owner, item)
		assert.NoError(err)
		testBodyLogRental := model.RequestPostLogsBody{
			OwnerID: owner.UserID,
			Type:    0,
			Purpose: "GetItemのテストのPurposeですrental1",
			DueDate: "2000-02-16",
			Count:   1,
		}

		reqBody, _ := json.Marshal(testBodyLogRental)
		req := httptest.NewRequest(echo.POST, "/api/items/"+strconv.Itoa(int(item.ID))+"/logs", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusCreated, rec.Code)

		req = httptest.NewRequest(echo.GET, "/api/items?rental=testUser", nil)
		req.Header.Set("Content-Type", "application/json")
		rec = httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusOK, rec.Code)
		gotItemsSuccess1 := []model.Item{}
		_ = json.NewDecoder(rec.Body).Decode(&gotItemsSuccess1)
		assert.NotEmpty(gotItemsSuccess1)
		exist1 := false
		for _, gotItem := range gotItemsSuccess1 {
			if item.Name == gotItem.Name {
				exist1 = true
			}
		}
		assert.Equal(true, exist1)

		req = httptest.NewRequest(echo.GET, "/api/items?user=testGetItemsUser", nil)
		req.Header.Set("Content-Type", "application/json")
		rec = httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusOK, rec.Code)
		gotItemsSuccess2 := []model.Item{}
		_ = json.NewDecoder(rec.Body).Decode(&gotItemsSuccess2)
		assert.NotEmpty(gotItemsSuccess2)
		exist2 := false
		for _, gotItem := range gotItemsSuccess2 {
			if item.Name == gotItem.Name {
				exist2 = true
			}
		}
		assert.Equal(true, exist2)
	})

}

func TestPostOwners(t *testing.T) {
	testBodyTrap := model.Item{
		Name:        "testPostOwnersTrapItem",
		Type:        1,
		Code:        "1920093013001",
		Description: "これは備品のテストです",
		ImgURL:      "http://example.com/testTrap.jpg",
	}

	testBodyKojin := model.Item{
		Name:        "testPostOwnersKojinItem",
		Type:        0,
		Code:        "9784049123945",
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
		assert.Equal(trap.ID, item.Owners[0].UserID)
	})

	t.Run("not admin user", func(t *testing.T) {
		assert := assert.New(t)
		e := echoSetupWithUser()

		user := model.User{
			Name:        "testUser",
			DisplayName: "テストユーザー",
			Admin:       false,
		}
		testUser, err := model.GetUserByName(user.Name)
		assert.NotEmpty(testUser)
		assert.NoError(err)

		testOwnerKojin := model.RequestPostOwnersBody{
			UserID:     int(testUser.ID),
			Rentalable: true,
			Count:      1,
		}

		bihin, _ := model.GetItemByName("testPostOwnersTrapItem")
		reqBody, _ := json.Marshal(testOwnerKojin)
		req := httptest.NewRequest(echo.POST, "/api/items/"+strconv.Itoa(int(bihin.ID))+"/owners", bytes.NewReader(reqBody))
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

		paramID := strconv.Itoa(int(item.ID))
		targetAPI := "/api/items/" + paramID + "/owners"

		reqBody, _ = json.Marshal(testOwnerKojin)
		req = httptest.NewRequest(echo.POST, targetAPI, bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec = httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusOK, rec.Code)

		_ = json.NewDecoder(rec.Body).Decode(&item)

		assert.Equal(testBodyKojin.Name, item.Name)
		assert.Equal(testUser.ID, item.Owners[0].UserID)
	})
}

func TestPostLikes(t *testing.T) {
	item, _ := model.CreateItem(model.Item{Name: "testPostLikesItem"})

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)
		e := echoSetupWithAdminUser()

		req := httptest.NewRequest(echo.POST, "/api/items/"+strconv.Itoa(int(item.ID))+"/likes", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusCreated, rec.Code)
	})

	t.Run("failuer", func(t *testing.T) {
		assert := assert.New(t)
		e := echoSetupWithAdminUser()

		req := httptest.NewRequest(echo.POST, "/api/items/"+strconv.Itoa(int(item.ID))+"/likes", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		req = httptest.NewRequest(echo.POST, "/api/items/"+strconv.Itoa(int(item.ID))+"/likes", nil)
		req.Header.Set("Content-Type", "application/json")
		rec = httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusBadRequest, rec.Code)
	})
}

func TestDeleteLikes(t *testing.T) {
	item, _ := model.CreateItem(model.Item{Name: "testDeleteLikesItem"})

	t.Run("failuer", func(t *testing.T) {
		assert := assert.New(t)
		e := echoSetupWithAdminUser()

		req := httptest.NewRequest(echo.DELETE, "/api/items/"+strconv.Itoa(int(item.ID))+"/likes", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusBadRequest, rec.Code)
	})

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)
		e := echoSetupWithAdminUser()

		req := httptest.NewRequest(echo.POST, "/api/items/"+strconv.Itoa(int(item.ID))+"/likes", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		req = httptest.NewRequest(echo.DELETE, "/api/items/"+strconv.Itoa(int(item.ID))+"/likes", nil)
		req.Header.Set("Content-Type", "application/json")
		rec = httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusCreated, rec.Code)
	})
}
