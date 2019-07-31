package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagTableName(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "tags", (&Tag{}).TableName())
}
