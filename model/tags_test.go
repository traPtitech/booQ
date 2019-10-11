package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagTableName(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "tags", (&Tag{}).TableName())
}

func TestCreateTag(t *testing.T) {
	t.Parallel()

	t.Run("failures", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		tag, err := CreateTag("")
		assert.Error(err)
		assert.Empty(tag)
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		tag, err := CreateTag("testCreateTag")
		assert.NoError(err)
		assert.NotEmpty(tag)
		assert.Equal(tag.Name, "testCreateTag")
	})
}

func TestAttachTag(t *testing.T) {
	t.Parallel()

	t.Run("failures", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		tag, err := AttachTag(99, 99)
		assert.Error(err)
		assert.Empty(tag)
		tag, err = CreateTag("testAttachTagFailTag")
		assert.NoError(err)
		tag, err = AttachTag(tag.ID, 99)
		assert.Error(err)
		assert.Empty(tag)
		item, err := CreateItem(Item{Name: "testAttachTagFailItem"})
		assert.NoError(err)
		tag, err = AttachTag(99, item.ID)
		assert.Error(err)
		assert.Empty(tag)
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		tag, err := CreateTag("testAttachTagSuccessTag")
		assert.NoError(err)
		item, err := CreateItem(Item{Name: "testAttachTagSuccessItem"})
		assert.NoError(err)
		tag, err = AttachTag(tag.ID, item.ID)
		assert.NoError(err)
		assert.NotEmpty(tag)
		assert.Equal(tag.Items[0].Name, "testAttachTagSuccessItem")
	})
}

func TestRemoveTag(t *testing.T) {
	t.Parallel()

	t.Run("failures", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		tag, err := RemoveTag(Tag{}, 99)
		assert.Error(err)
		assert.Empty(tag)
		createdTag, err := CreateTag("testRemoveTagFailTag")
		assert.NoError(err)
		tag, err = RemoveTag(createdTag, 99)
		assert.Error(err)
		assert.Empty(tag)
		item, err := CreateItem(Item{Name: "testRemoveTagFailItem"})
		assert.NoError(err)
		tag, err = RemoveTag(Tag{}, item.ID)
		assert.Error(err)
		assert.Empty(tag)
		tag, err = RemoveTag(createdTag, item.ID)
		assert.Error(err)
		assert.Empty(tag)
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		tag, err := CreateTag("testRemoveTagSuccessTag")
		assert.NoError(err)
		item, err := CreateItem(Item{Name: "testRemoveTagSuccessItem"})
		assert.NoError(err)
		tag, err = AttachTag(tag.ID, item.ID)
		assert.NoError(err)
		tag, err = RemoveTag(tag, item.ID)
		assert.NoError(err)
		assert.NotEmpty(tag)
		exist := false
		for _, tagItem := range tag.Items {
			if tagItem.ID == item.ID {
				exist = true
			}
		}
		assert.Equal(false, exist)
	})
}

func TestGetTagByName(t *testing.T) {
	t.Parallel()

	t.Run("failures", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		tag, err := GetTagByName("testGetTagByNameTag")
		assert.Error(err)
		assert.Empty(tag)
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		tag, err := CreateTag("testGetTagByNameTag")
		assert.NoError(err)
		item, err := CreateItem(Item{Name: "testGetTagByNameItem"})
		assert.NoError(err)
		tag, err = AttachTag(tag.ID, item.ID)
		assert.NoError(err)
		tag, err = GetTagByName(tag.Name)
		assert.NoError(err)
		assert.NotEmpty(tag)
		assert.Equal("testGetTagByNameTag", tag.Name)
		exist := false
		for _, tagItem := range tag.Items {
			if tagItem.ID == item.ID {
				exist = true
				assert.Equal("testGetTagByNameItem", item.Name)
			}
		}
		assert.Equal(true, exist)
	})
}

func TestGetTagByID(t *testing.T) {
	t.Parallel()

	t.Run("failures", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		tag, err := GetTagByID(99)
		assert.Error(err)
		assert.Empty(tag)
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		tag, err := CreateTag("testGetTagByIDTag")
		assert.NoError(err)
		item, err := CreateItem(Item{Name: "testGetTagByIDItem"})
		assert.NoError(err)
		tag, err = AttachTag(tag.ID, item.ID)
		assert.NoError(err)
		tag, err = GetTagByID(tag.ID)
		assert.NoError(err)
		assert.NotEmpty(tag)
		assert.Equal("testGetTagByIDTag", tag.Name)
		exist := false
		for _, tagItem := range tag.Items {
			if tagItem.ID == item.ID {
				exist = true
				assert.Equal("testGetTagByIDItem", item.Name)
			}
		}
		assert.Equal(true, exist)
	})
}
