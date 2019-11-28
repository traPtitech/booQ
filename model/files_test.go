package model

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestFile_TableName(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "files", (&File{}).TableName())
}

func TestCreateFile(t *testing.T) {
	t.Parallel()

	user, err := CreateUser(User{Name: "testCreateFileUser"})
	require.NoError(t, err)

	f, err := CreateFile(user.ID, strings.NewReader("test file"), "txt")
	if assert.NoError(t, err) {
		assert.NotEmpty(t, f.ID)
		assert.EqualValues(t, user.ID, f.UploadUserID)
	}
}
