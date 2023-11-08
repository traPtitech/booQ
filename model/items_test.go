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

		_, _ = CreateItem(Item{Name: "testCreateItemDuplicateFail", Type: 1})
		item, err = CreateItem(Item{Name: "testCreateItemDuplicateFail", Type: 1})
		assert.Error(err)
		assert.Empty(item)
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		item, err := CreateItem(Item{Name: "testCreateItemSuccess"})
		assert.NoError(err)
		assert.NotEmpty(item)

		_, _ = CreateItem(Item{Name: "testCreateItemDuplicateSuccess", Type: 1})
		item, err = CreateItem(Item{Name: "testCreateItemDuplicateSuccess", Type: 0})
		assert.NoError(err)
		assert.NotEmpty(item)
	})
}

func TestRegisterOwner(t *testing.T) {
	t.Run("make success", func(t *testing.T) {
		user, err := CreateUser(User{Name: "testRegisterOwnerUser"})
		assert := assert.New(t)
		assert.NoError(err)
		var owner Owner
		owner.UserID = user.ID
		owner.Rentalable = true
		owner.Count = 1
		item, err := CreateItem(Item{Name: "testRegisterOwnerItem"})
		assert.NoError(err)
		item2, err := RegisterOwner(owner, item)
		assert.NoError(err)
		assert.NotEmpty(item2)
		assert.Equal(user.ID, item2.Owners[0].UserID)
		assert.Equal(user.Name, item2.Owners[0].User.Name)
	})
}

func TestAddOwner(t *testing.T) {
	user, _ := CreateUser(User{Name: "testAddOwnerUser"})
	var owner Owner
	owner.UserID = user.ID
	owner.Rentalable = true
	owner.Count = 5
	item, _ := CreateItem(Item{Name: "testAddOwnerItem"})
	t.Run("decreace fail", func(t *testing.T) {
		assert := assert.New(t)
		item1, err := AddOwner(owner, item)

		assert.Error(err)
		assert.Empty(item1)
	})
	t.Run("add & decreace success", func(t *testing.T) {
		assert := assert.New(t)

		item, err := RegisterOwner(owner, item)
		assert.NoError(err)
		assert.NotEmpty(item)

		owner.Count = 2
		item, err = AddOwner(owner, item)
		exist := false
		for _, owner := range item.Owners {
			if owner.UserID == user.ID {
				exist = true
				assert.Equal(2, owner.Count)
				assert.Equal(owner.User.Name, user.Name)
			}
		}
		assert.Equal(true, exist)
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

		// TODO: ちゃんとlikeしてるやつはIsLikedがtrueになってるかチェックする
		res, err := GetItems(0)
		assert.NoError(err)
		assert.NotEmpty(res)
		for _, value := range res {
			if value.Name == "testAllItemItem" {
				assert.Equal(user.ID, value.Owners[0].UserID)
				assert.Equal(user.Name, value.Owners[0].User.Name)
				break
			}
			continue
		}
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

func TestSearchItems(t *testing.T) {
	user, _ := CreateUser(User{Name: "testSearchItemUser"})
	var owner Owner
	owner.UserID = user.ID
	owner.Rentalable = true
	owner.Count = 1
	item, _ := CreateItem(Item{Name: "testSearchItemItem"})
	_, _ = CreateItem(Item{Name: "testSearchItemItem1"})
	_, _ = CreateItem(Item{Name: "testSearchItemsItem"})
	_, _ = RegisterOwner(owner, item)

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		// TODO: ちゃんとlikeしてるやつはIsLikedがtrueになってるかチェックする
		res, err := GetItems(0)
		assert.NoError(err)
		assert.NotEmpty(res)
		for _, value := range res {
			if value.Name == "testSearchItemItem" {
				assert.Equal(user.ID, value.Owners[0].UserID)
				assert.Equal(user.Name, value.Owners[0].User.Name)
				break
			}
			continue
		}

		items, err := SearchItems("SearchItemItem")
		assert.NoError(err)
		var existSearchItem = false
		var existSearchItem1 = false
		var existTestSearchItems = false

		for _, value := range items {
			if value.Name == "testSearchItemItem" {
				existSearchItem = true
			}
			if value.Name == "testSearchItemItem1" {
				existSearchItem1 = true
			}
			if value.Name == "testSearchItemsItem" {
				existTestSearchItems = true
			}
		}
		assert.Equal(true, existSearchItem)
		assert.Equal(true, existSearchItem1)
		assert.Equal(false, existTestSearchItems)
		assert.NotEmpty(items)
	})
}

func TestCreateLike(t *testing.T) {
	user, _ := CreateUser(User{Name: "testPushLikeUser"})
	item, _ := CreateItem(Item{Name: "testPushLikeItem"})

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)
		item, err := CreateLike(item.ID, user.ID)

		assert.Equal(user.ID, item.Likes[0].ID)
		assert.Equal(user.Name, item.Likes[0].Name)
		assert.NoError(err)
		assert.NotEmpty(item)
	})

	t.Run("failer", func(t *testing.T) {
		assert := assert.New(t)
		item, err := CreateLike(item.ID, user.ID)
		t.Log(item.Likes)

		assert.Error(err)
		assert.Empty(item)
	})
}

func TestCancelLike(t *testing.T) {
	user, _ := CreateUser(User{Name: "testDeleteLikeUser"})
	item, _ := CreateItem(Item{Name: "testDeleteLikeItem"})

	t.Run("failer", func(t *testing.T) {
		assert := assert.New(t)
		item, err := CancelLike(item.ID, user.ID)

		assert.Error(err)
		assert.Empty(item)
	})

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)
		item, err := CreateLike(item.ID, user.ID)
		assert.NotEmpty(item)
		assert.NoError(err)
		assert.Equal(user.ID, item.Likes[0].ID)
		assert.Equal(user.Name, item.Likes[0].Name)
		item, err = CancelLike(item.ID, user.ID)
		assert.NotEmpty(item)
		assert.NoError(err)

		exist := false
		for _, likeUser := range item.Likes {
			if likeUser.Name == user.Name {
				exist = true
			}
		}
		assert.Equal(false, exist)
	})
}

func TestDestroyItem(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)
		item, err := CreateItem(Item{Name: "testDestroyItemSuccess"})
		assert.NotEmpty(item)
		assert.NoError(err)
		updateItem, err := DestroyItem(item)
		assert.NotEmpty(updateItem)
		assert.NoError(err)
		item, err = GetItemByID(item.ID)
		assert.Empty(item)
		assert.Error(err)
	})
}

func TestUpdateItem(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)
		item, err := CreateItem(Item{Name: "testUpdateItemSuccess", Type: 1})
		assert.NotEmpty(item)
		assert.NoError(err)
		body := RequestPutItemBody{}
		body.Name = "updateTestUpdateItemSuccess"
		updateItem, err := UpdateItem(&item, &body, false)
		assert.NotEmpty(updateItem)
		assert.NoError(err)
		assert.Equal("updateTestUpdateItemSuccess", updateItem.Name)
		assert.Equal(1, updateItem.Type)
	})
}

func TestRentalItem(t *testing.T) {
	t.Parallel()

	t.Run("failuer", func(t *testing.T) {
		assert := assert.New(t)
		user, err := CreateUser(User{Name: "testRentalItemFailUser"})
		assert.NotEmpty(user)
		assert.NoError(err)
		owner, err := CreateUser(User{Name: "testRentalItemFailOwner"})
		assert.NotEmpty(owner)
		assert.NoError(err)
		item, err := CreateItem(Item{Name: "testRentalItemFailItem"})
		assert.NotEmpty(item)
		assert.NoError(err)
		rentalUserReturn := RentalUser{
			UserID:  user.ID,
			Count:   2,
			OwnerID: owner.ID,
		}
		failItem1, err := RentalItem(rentalUserReturn, item)
		assert.Empty(failItem1)
		assert.Error(err)
		rentalUserRental := RentalUser{
			UserID:  user.ID,
			Count:   -1,
			OwnerID: owner.ID,
		}
		successItem, err := RentalItem(rentalUserRental, item)
		assert.NotEmpty(successItem)
		assert.NoError(err)
		failItem2, err := RentalItem(rentalUserReturn, item)
		assert.Empty(failItem2)
		assert.Error(err)
	})

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)
		user, err := CreateUser(User{Name: "testRentalItemSuccessUser"})
		assert.NotEmpty(user)
		assert.NoError(err)
		owner, err := CreateUser(User{Name: "testRentalItemSuccessOwner"})
		assert.NotEmpty(owner)
		assert.NoError(err)
		item, err := CreateItem(Item{Name: "testRentalItemSuccessItem"})
		assert.NotEmpty(item)
		assert.NoError(err)
		rentalUserRental := RentalUser{
			UserID:  user.ID,
			Count:   -2,
			OwnerID: owner.ID,
		}
		successItem1, err := RentalItem(rentalUserRental, item)
		assert.NotEmpty(successItem1)
		assert.NoError(err)
		exist1 := false
		for _, rentalUser := range successItem1.RentalUsers {
			if rentalUser.UserID == user.ID {
				exist1 = true
			}
		}
		assert.Equal(true, exist1)
		rentalUserReturn := RentalUser{
			UserID:  user.ID,
			Count:   1,
			OwnerID: owner.ID,
		}
		successItem2, err := RentalItem(rentalUserReturn, item)
		assert.NotEmpty(successItem2)
		assert.NoError(err)
		assert.Equal(-1, successItem2.RentalUsers[0].Count)
		successItem3, err := RentalItem(rentalUserReturn, item)
		assert.NotEmpty(successItem3)
		assert.NoError(err)
	})
}

func TestSearchItemByRental(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)
		user, err := CreateUser(User{Name: "testSearchItemByRentalUser"})
		assert.NotEmpty(user)
		assert.NoError(err)
		owner, err := CreateUser(User{Name: "testSearchItemByRentalOwner"})
		assert.NotEmpty(owner)
		assert.NoError(err)
		item1, err := CreateItem(Item{Name: "testSearchItemByRentalItem1"})
		assert.NotEmpty(item1)
		assert.NoError(err)
		item2, err := CreateItem(Item{Name: "testSearchItemByRentalItem2"})
		assert.NotEmpty(item2)
		assert.NoError(err)
		item3, err := CreateItem(Item{Name: "testSearchItemByRentalItem3"})
		assert.NotEmpty(item3)
		assert.NoError(err)
		rentalUserRental := RentalUser{
			UserID:  user.ID,
			Count:   -1,
			OwnerID: owner.ID,
		}
		successItem1, err := RentalItem(rentalUserRental, item1)
		assert.NotEmpty(successItem1)
		assert.NoError(err)
		successItem2, err := RentalItem(rentalUserRental, item2)
		assert.NotEmpty(successItem2)
		assert.NoError(err)
		rentalUserReturn := RentalUser{
			UserID:  user.ID,
			Count:   1,
			OwnerID: owner.ID,
		}
		successItem2, err = RentalItem(rentalUserReturn, item2)
		assert.NotEmpty(successItem2)
		assert.NoError(err)
		// TODO: ちゃんとlikeしてるやつはIsLikedがtrueになってるかチェックする
		res, err := SearchItemByRental(user.ID, 0)
		assert.NotEmpty(successItem2)
		assert.NoError(err)
		exist1 := false
		exist2 := false
		exist3 := false
		for _, value := range res {
			if value.Name == item1.Name {
				exist1 = true
			}
			if value.Name == item2.Name {
				exist2 = true
			}
			if value.Name == item3.Name {
				exist3 = true
			}
		}
		assert.Equal(true, exist1)
		assert.Equal(false, exist2)
		assert.Equal(false, exist3)
	})
}

func TestSearchItemByOwner(t *testing.T) {
	t.Run("failuer", func(t *testing.T) {
		assert := assert.New(t)
		// TODO: ちゃんとlikeしてるやつはIsLikedがtrueになってるかチェックする
		items, err := SearchItemByOwner("testSearchItemByOwnerFailOwner", 0)
		assert.Empty(items)
		assert.Error(err)
	})

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)
		ownerUser, err := CreateUser(User{Name: "testSearchItemByOwSucOwner"})
		assert.NotEmpty(ownerUser)
		assert.NoError(err)
		item1, err := CreateItem(Item{Name: "testSearchItemByOwSucItem1"})
		assert.NotEmpty(item1)
		assert.NoError(err)
		item2, err := CreateItem(Item{Name: "testSearchItemByOwSucItem2"})
		assert.NotEmpty(item2)
		assert.NoError(err)
		owner := Owner{
			UserID:     ownerUser.ID,
			User:       ownerUser,
			Count:      1,
			Rentalable: true,
		}
		item1, err = RegisterOwner(owner, item1)
		assert.NotEmpty(item1)
		assert.NoError(err)
		// TODO: ちゃんとlikeしてるやつはIsLikedがtrueになってるかチェックする
		res, err := SearchItemByOwner(ownerUser.Name, 0)
		assert.NotEmpty(res)
		assert.NoError(err)
		exist1 := false
		exist2 := false
		for _, value := range res {
			if value.Name == item1.Name {
				for _, nowOwner := range value.Owners {
					if nowOwner.User.Name == ownerUser.Name {
						exist1 = true
					}
				}
			}
			if value.Name == item2.Name {
				for _, nowOwner := range value.Owners {
					if nowOwner.User.Name == ownerUser.Name {
						exist2 = true
					}
				}
			}
		}
		assert.Equal(true, exist1)
		assert.Equal(false, exist2)
	})
}
