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

		item, err = CreateItem(Item{Name: "testCreateItemFail1", Code: "1234567891012"})
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

func TestRegisterOwner(t *testing.T) {
	user, _ := CreateUser(User{Name: "testRegisterOwnerUser"})
	var owner Owner
	owner.OwnerID = user.ID
	owner.Rentalable = true
	owner.Count = 1
	item, _ := CreateItem(Item{Name: "testRegisterOwnerItem"})
	item2, err := RegisterOwner(owner, item)
	t.Run("make success", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		assert.Equal(user.ID, item2.Owners[0].OwnerID)
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
	user, _ := CreateUser(User{Name: "testAllItemUser"})
	var owner Owner
	owner.OwnerID = user.ID
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
				assert.Equal(user.ID, value.Owners[0].OwnerID)
				break
			}
			continue
		}
		assert.NoError(err)
		assert.NotEmpty(items)
	})
}

func TestGetItemByID(t *testing.T) {

	user1, _ := CreateUser(User{Name: "testGetItemByIDOwner"})
	user2, _ := CreateUser(User{Name: "testGetItemByIDUser"})
	owner := Owner{
		OwnerID:    user1.ID,
		Rentalable: true,
		Count:      1,
	}

	t.Run("success", func(t *testing.T) {

		t.Parallel()
		assert := assert.New(t)
		item, err := CreateItem(Item{Name: "testGetItemItem"})
		assert.NoError(err)
		_, err = RegisterOwner(owner, item)
		assert.NoError(err)
		_, err = CreateLog(Log{ItemID: item.ID, OwnerID: owner.OwnerID, UserID: user2.ID, Type: 0, Count: 1})
		assert.NoError(err)

		gotItem, err := GetItemByID(item.ID)

		assert.NoError(err)
		assert.NotEmpty(gotItem)
		assert.Equal(gotItem.Name, "testGetItemItem")
		assert.Equal(gotItem.Owners[0].OwnerID, user1.ID)
		assert.Equal(gotItem.Logs[0].OwnerID, user1.ID)
		assert.Equal(gotItem.Logs[0].Count, 1)
		assert.Equal(gotItem.Logs[0].ItemID, item.ID)
	})
}
