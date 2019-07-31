package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemTableName(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "items", (&Item{}).TableName())
}
