package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLikemapTableName(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "likemaps", (&Likemap{}).TableName())
}
