package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentTableName(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "comments", (&Comment{}).TableName())
}
