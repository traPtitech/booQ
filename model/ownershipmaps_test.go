package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOwnershipmapTableName(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "ownershipmaps", (&Ownershipmap{}).TableName())
}

func TestCreateOwnershimap(t *testing.T) {
	t.Parallel()

	t.Run("failures", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		ownershipmap, err := CreateOwnershipmap(Ownershipmap{})
		assert.Error(err)
		assert.Empty(ownershipmap)
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		ownershipmap, err := CreateOwnershipmap(Ownershipmap{ItemID: 3, UserID: 1})
		assert.NoError(err)
		assert.NotEmpty(ownershipmap)
	})
}
