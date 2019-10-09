package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentTableName(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "comments", (&Comment{}).TableName())
}

func TestCreateComment(t *testing.T) {
	t.Parallel()
	item, _ := CreateItem(Item{Name: "testCreateComment"})

	t.Run("failures", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		comment, err := CreateComment(Comment{})
		assert.Error(err)
		assert.Empty(comment)
		comment, err = CreateComment(Comment{ItemID: 999})
		assert.Error(err)
		assert.Empty(comment)
		comment, err = CreateComment(Comment{ItemID: item.ID})
		assert.Error(err)
		assert.Empty(comment)
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		comment, err := CreateComment(Comment{ItemID: item.ID, Text: "testCreateCommentText"})
		assert.NoError(err)
		assert.NotEmpty(comment)
	})
}
