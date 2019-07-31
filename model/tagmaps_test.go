package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagmapTableName(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "tagmaps", (&Tagmap{}).TableName())
}
