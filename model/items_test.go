package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemTableName(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "items", (&Item{}).TableName())
}

func TestCreateItem(t *testing.T) {
	t.Parallel()

	t.Run("failures", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		item, err := CreateItem(Item{})
		assert.Error(err)
		assert.Empty(item)

		_, _ = CreateItem(Item{Name: "testCreateItemFail", Code: "1234567891012"})
		item, err = CreateItem(Item{Name: "testCreateItemFail"})
		assert.Error(err)
		assert.Empty(item)

		item, err = CreateItem(Item{Code: "1234567891012"})
		assert.Error(err)
		assert.Empty(item)
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		item, err := CreateItem(Item{Name: "testCreateItemSuccess"})
		assert.NoError(err)
		assert.NotEmpty(item)
	})
}

func TestGetItems(t *testing.T) {
	t.Parallel()

	user, _ := CreateUser(User{Name: "testAllItemUser"})
	var owner Owner
	owner.Owner = user
	owner.Rentalable = true
	item, _ := CreateItem(Item{Name: "testAllItemItem"})
	_, _ = RegisterOwner(owner, item)
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		items, err := GetItems()
		for _, value := range items {
			if value.Name == "testAllItemItem" {
				assert.Equal("testAllItemUser", value.Owners[0].Owner.Name)
				break
			}
			continue
		}
		assert.NoError(err)
		assert.NotEmpty(items)
	})
}
