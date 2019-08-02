package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserTableName(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "users", (&User{}).TableName())
}

func TestGetUser(t *testing.T) {
	t.Parallel()

	t.Run("failures", func(t *testing.T) {
		assert := assert.New(t)

		user, err := GetUser(User{})
		assert.Error(err)
		assert.Empty(user)

		user, err = GetUser(User{Name: "nothing user"})
		assert.NoError(err)
		assert.Empty(user)
	})

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)

		user, err := GetUser(User{Name: "traP"})
		assert.NoError(err)
		assert.NotEmpty(user)
		assert.Equal("traP", user.Name)
	})
}

func TestCreateUser(t *testing.T) {
	t.Parallel()

	t.Run("failures", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		user, err := CreateUser(User{})
		assert.Error(err)
		assert.Empty(user)
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		user, err := CreateUser(User{Name: "test"})
		assert.NoError(err)
		assert.NotEmpty(user)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Parallel()

	t.Run("failures", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		user, err := UpdateUser(User{IconFileID:"testfile"})
		assert.Error(err)
		assert.Empty(user)
	})

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)

		user1, err1 := CreateUser(User{Name: "test3"})
		assert.NoError(err1)
		assert.NotEmpty(user1)

		user, err := UpdateUser(User{Name:"test3",IconFileID:"testfile"})
		assert.NoError(err)
		assert.NotEmpty(user)
		assert.Equal("testfile", user.IconFileID)
	})
}
