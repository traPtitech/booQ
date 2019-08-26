package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOwnershipmapTableName(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "ownershipmaps", (&Ownershipmap{}).TableName())
}
