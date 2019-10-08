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
	owner.UserID = user.ID
	owner.Rentalable = true
	owner.Count = 1
	item, _ := CreateItem(Item{Name: "testRegisterOwnerItem"})
	item2, err := RegisterOwner(owner, item)

	t.Run("make success", func(t *testing.T) {
		assert := assert.New(t)

		assert.Equal(user.ID, item2.Owners[0].UserID)
		assert.Equal(user.Name, item2.Owners[0].User.Name)
		assert.NoError(err)
		assert.NotEmpty(item2)
	})

	t.Run("add success", func(t *testing.T) {
		assert := assert.New(t)

		owner.Count = 5
		item, err := RegisterOwner(owner, item)

		assert.Equal(6, item.Owners[0].Count)
		assert.Equal(item.Owners[0].User.Name, user.Name)

		assert.NoError(err)
		assert.NotEmpty(item)
	})
}

func TestGetItems(t *testing.T) {
	user, _ := CreateUser(User{Name: "testAllItemUser"})
	var owner Owner
	owner.UserID = user.ID
	owner.Rentalable = true
	owner.Count = 1
	item, _ := CreateItem(Item{Name: "testAllItemItem"})
	_, _ = RegisterOwner(owner, item)

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		items, err := GetItems()
		assert.NoError(err)

		for _, value := range items {
			if value.Name == "testAllItemItem" {
				assert.Equal(user.ID, value.Owners[0].UserID)
				assert.Equal(user.Name, value.Owners[0].User.Name)
				break
			}
			continue
		}
		assert.NotEmpty(items)
	})
}

func TestGetItemByID(t *testing.T) {
	ownerUser, _ := CreateUser(User{Name: "testGetItemByIDOwner"})
	rentalUser, _ := CreateUser(User{Name: "testGetItemByIDUser"})
	owner := Owner{
		UserID:     ownerUser.ID,
		Rentalable: true,
		Count:      1,
	}

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)
		item, err := CreateItem(Item{Name: "testGetItemItem"})
		assert.NoError(err)
		_, err = RegisterOwner(owner, item)
		assert.NoError(err)
		_, err = CreateLog(Log{ItemID: item.ID, OwnerID: owner.UserID, UserID: rentalUser.ID, Type: 0, Count: 1})
		assert.NoError(err)

		gotItem, err := GetItemByID(item.ID)

		assert.NoError(err)
		assert.NotEmpty(gotItem)
		assert.Equal(gotItem.Name, "testGetItemItem")
		assert.Equal(gotItem.Owners[0].UserID, ownerUser.ID)
		assert.Equal(gotItem.Owners[0].User.Name, ownerUser.Name)
		assert.Equal(gotItem.Logs[0].OwnerID, ownerUser.ID)
		assert.Equal(gotItem.Logs[0].Owner.Name, ownerUser.Name)
		assert.Equal(gotItem.Logs[0].Count, 1)
		assert.Equal(gotItem.Logs[0].ItemID, item.ID)
	})
}

func TestGetItemByName(t *testing.T) {
	ownerUser, _ := CreateUser(User{Name: "testGetItemByNameOwner"})
	rentalUser, _ := CreateUser(User{Name: "testGetItemByNameUser"})
	owner := Owner{
		UserID:     ownerUser.ID,
		Rentalable: true,
		Count:      1,
	}

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)

		item, err := CreateItem(Item{Name: "testGetItemByNameItem"})
		assert.NoError(err)
		_, err = RegisterOwner(owner, item)
		assert.NoError(err)
		_, err = CreateLog(Log{ItemID: item.ID, OwnerID: owner.UserID, UserID: rentalUser.ID, Type: 0, Count: 1})
		assert.NoError(err)
		gotItem, err := GetItemByName(item.Name)

		assert.NoError(err)
		assert.NotEmpty(gotItem)
		assert.Equal(gotItem.Name, "testGetItemByNameItem")
		assert.Equal(gotItem.Owners[0].UserID, ownerUser.ID)
		assert.Equal(gotItem.Owners[0].User.Name, ownerUser.Name)
		assert.Equal(gotItem.Logs[0].OwnerID, ownerUser.ID)
		assert.Equal(gotItem.Logs[0].Owner.Name, ownerUser.Name)
		assert.Equal(gotItem.Logs[0].Count, 1)
		assert.Equal(gotItem.Logs[0].ItemID, item.ID)
	})
}

func TestPushLike(t *testing.T) {
	user, _ := CreateUser(User{Name: "testPushLikeUser"})
	item, _ := CreateItem(Item{Name: "testPushLikeItem"})

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)
		item, err := PushLike(item.ID, user.ID)

		assert.Equal(user.ID, item.Likes[0].ID)
		assert.Equal(user.Name, item.Likes[0].Name)
		assert.NoError(err)
		assert.NotEmpty(item)
	})

	t.Run("failer", func(t *testing.T) {
		assert := assert.New(t)
		item, err := PushLike(item.ID, user.ID)
		t.Log(item.Likes)

		assert.Error(err)
		assert.Empty(item)
	})
}
