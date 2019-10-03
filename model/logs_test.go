package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogTableName(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "logs", (&Log{}).TableName())
}
