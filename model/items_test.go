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
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		item, err := CreateItem(Item{Name: "test"})
		assert.NoError(err)
		assert.NotEmpty(item)
	})
}

func TestRegisterOwner(t *testing.T) {
	t.Parallel()

	user, _ := CreateUser(User{Name: "testRegisterOwnerUser"})
	var owner Owner
	owner.OwnerID = int(user.ID)
	owner.Rentalable = true
	owner.Count = 1
	item, _ := CreateItem(Item{Name: "testRegisterOwnerItem"})
	item2, err := RegisterOwner(owner, item)
	t.Run("make success", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		assert.Equal(int(user.ID), item2.Owners[0].OwnerID)
		assert.NoError(err)
		assert.NotEmpty(item2)
	})

	t.Run("add success", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		owner.Count = 5
		item, err := RegisterOwner(owner, item)

		assert.Equal(6, item.Owners[0].Count)

		assert.NoError(err)
		assert.NotEmpty(item)
	})
}

func TestGetItems(t *testing.T) {
	t.Parallel()

	user, _ := CreateUser(User{Name: "testAllItemUser"})
	var owner Owner
	owner.OwnerID = int(user.ID)
	owner.Rentalable = true
	owner.Count = 1
	item, _ := CreateItem(Item{Name: "testAllItemItem"})
	_, _ = RegisterOwner(owner, item)
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		items, err := GetItems()
		for _, value := range items {
			if value.Name == "testAllItemItem" {
				assert.Equal(int(user.ID), value.Owners[0].OwnerID)
				break
			}
			continue
		}
		assert.NoError(err)
		assert.NotEmpty(items)
	})
}
