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
