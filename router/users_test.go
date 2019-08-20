package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"github.com/traPtitech/booQ/model"
)

// TestGetUsersMe /users/me のテスト
func TestGetUsersMe(t *testing.T) {
	t.Parallel()

	t.Run("admin user", func(t *testing.T) {
		assert := assert.New(t)

		adminUser := model.User{
			Name:        "traP",
			DisplayName: "traP",
			IconFileID:  "099eed74-3ab3-4655-ac37-bc7df1139b3d",
			Admin:       true,
		}

		e := echoSetupWithAdminUser()

		req := httptest.NewRequest(echo.GET, "/api/users/me", nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusOK, rec.Code)

		body, _ := ioutil.ReadAll(rec.Body)
		user := model.User{}
		_ = json.Unmarshal(body, &user)

		assert.Equal(user.Name, adminUser.Name)
		assert.Equal(user.DisplayName, adminUser.DisplayName)
		assert.Equal(user.IconFileID, adminUser.IconFileID)
		assert.Equal(user.Admin, adminUser.Admin)
	})

	t.Run("new user", func(t *testing.T) {
		assert := assert.New(t)

		testUser := model.User{
			Name:        "testUser",
			DisplayName: "テストユーザー",
			IconFileID:  "099eed74-3ab3-4655-ac37-bc7df1139b3d",
			Admin:       false,
		}

		e := echoSetupWithUser()

		req := httptest.NewRequest(echo.GET, "/api/users/me", nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusOK, rec.Code)

		body, _ := ioutil.ReadAll(rec.Body)
		user := model.User{}
		_ = json.Unmarshal(body, &user)

		assert.Equal(user.Name, testUser.Name)
		assert.Equal(user.DisplayName, testUser.DisplayName)
		assert.Equal(user.IconFileID, testUser.IconFileID)
		assert.Equal(user.Admin, testUser.Admin)
	})
}
