package router

import (
	"bytes"
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

// TestGetUserMe PUT /users のテスト
func TestPutUsers(t *testing.T) {
	t.Parallel()

	testUser := model.User{
		Name:        "PutUser",
		DisplayName: "テストユーザー",
		IconFileID:  "099eed74-3ab3-4655-ac37-bc7df1139b3d",
		Admin:       false,
	}
	_, _ = model.CreateUser(testUser)

	testBody := model.User{
		Name:        "PutUser",
		DisplayName: "変更されたテストユーザー",
		IconFileID:  "099eed74-3ab3-4655-ac37-testtesttest",
		Admin:       true,
	}

	t.Run("not admin user", func(t *testing.T) {
		assert := assert.New(t)
		e := echoSetupWithUser()

		reqBody, _ := json.Marshal(testBody)
		req := httptest.NewRequest(echo.PUT, "/api/users", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusForbidden, rec.Code)

		user, err := model.GetUserByName(testUser.Name)
		assert.NoError(err)
		assert.Equal(user.DisplayName, testUser.DisplayName)
		assert.Equal(user.IconFileID, testUser.IconFileID)
	})

	t.Run("admin user", func(t *testing.T) {
		assert := assert.New(t)
		e := echoSetupWithAdminUser()

		reqBody, _ := json.Marshal(testBody)
		req := httptest.NewRequest(echo.PUT, "/api/users", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusOK, rec.Code)

		user := model.User{}
		_ = json.NewDecoder(rec.Body).Decode(&user)

		assert.Equal(testBody.Name, user.Name)
		assert.Equal(testBody.DisplayName, user.DisplayName)
		assert.Equal(testBody.IconFileID, user.IconFileID)
		assert.Equal(testBody.Admin, user.Admin)
	})
}
