package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogTableName(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "logs", (&Log{}).TableName())
}

func TestCreateLog(t *testing.T) {
	t.Parallel()

	t.Run("failures", func(t *testing.T) {
		assert := assert.New(t)

		log, err := CreateLog(Log{})
		assert.Error(err)
		assert.Empty(log)

		log, err = CreateLog(Log{ItemID: 66, OwnerID: 66})
		assert.Error(err)
		assert.Empty(log)
	})

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)

		owner, _ := GetUserByName("traP")
		item, _ := CreateItem(Item{Name: "testItemForCreateLog"})

		log, err := CreateLog(Log{ItemID: item.ID, OwnerID: owner.ID, Type: 0})
		assert.NoError(err)
		assert.NotEmpty(log)
		assert.Equal(owner.ID, log.OwnerID)
		assert.Equal(item.ID, log.ItemID)
	})
}

func TestGetLatestLogs(t *testing.T) {
	assert := assert.New(t)
	item, err := CreateItem(Item{Name: "testGetLatestLogsItem"})
	assert.NoError(err)
	itemID := item.ID
	user1, err := CreateUser(User{Name: "testGetUserLatestLogsUser1"})
	assert.NoError(err)
	user2, err := CreateUser(User{Name: "testGetUserLatestLogsUser2"})
	assert.NoError(err)
	ownerUser1, err := GetUserByName("traP")
	assert.NoError(err)
	owner1 := Owner{
		UserID:     ownerUser1.ID,
		Rentalable: true,
		Count:      1,
	}
	ownerUser2, _ := GetUserByName("sienka")
	owner2 := Owner{
		UserID:     ownerUser2.ID,
		Rentalable: true,
		Count:      1,
	}

	t.Run("success", func(t *testing.T) {
		_, err := RegisterOwner(owner1, item)
		assert.NoError(err)
		_, err = RegisterOwner(owner2, item)
		assert.NoError(err)
		_, err = CreateLog(Log{ItemID: itemID, UserID: user1.ID, OwnerID: ownerUser1.ID, Type: 0, Count: 0})
		assert.NoError(err)
		_, err = CreateLog(Log{ItemID: itemID, UserID: user2.ID, OwnerID: ownerUser1.ID, Type: 0, Count: 1})
		assert.NoError(err)
		_, err = CreateLog(Log{ItemID: itemID, UserID: user2.ID, OwnerID: ownerUser2.ID, Type: 0, Count: 0})
		assert.NoError(err)
		_, err = CreateLog(Log{ItemID: itemID, UserID: user1.ID, OwnerID: ownerUser2.ID, Type: 0, Count: 1})
		assert.NoError(err)

		item, err = GetItemByID(item.ID)
		assert.NoError(err)
		latestLogs, err := GetLatestLogs(item.Logs)
		assert.NoError(err)
		assert.NotEmpty(latestLogs)
		for _, log := range latestLogs {
			if log.OwnerID == ownerUser1.ID {
				assert.Equal(ownerUser1.Name, log.Owner.Name)
				assert.Equal(user2.ID, log.UserID)
				assert.Equal(user2.Name, log.User.Name)
				assert.Equal(1, log.Count)
			}
			if log.OwnerID == ownerUser2.ID {
				assert.Equal(ownerUser2.Name, log.Owner.Name)
				assert.Equal(user1.ID, log.UserID)
				assert.Equal(user1.Name, log.User.Name)
				assert.Equal(1, log.Count)
			}
		}
	})
}

func TestGetLatestLog(t *testing.T) {
	assert := assert.New(t)
	item, err := CreateItem(Item{Name: "testGetLatestLogItem"})
	assert.NoError(err)
	itemID := item.ID
	user1, err := CreateUser(User{Name: "testGetUserLatestLogUser1"})
	assert.NoError(err)
	user2, err := CreateUser(User{Name: "testGetUserLatestLogUser2"})
	assert.NoError(err)
	ownerUser1, err := GetUserByName("traP")
	assert.NoError(err)
	owner1 := Owner{
		UserID:     ownerUser1.ID,
		Rentalable: true,
		Count:      1,
	}
	ownerUser2, _ := GetUserByName("sienka")
	owner2 := Owner{
		UserID:     ownerUser2.ID,
		Rentalable: true,
		Count:      1,
	}

	t.Run("success", func(t *testing.T) {
		_, err := RegisterOwner(owner1, item)
		assert.NoError(err)
		_, err = RegisterOwner(owner2, item)
		assert.NoError(err)
		_, err = CreateLog(Log{ItemID: itemID, UserID: user1.ID, OwnerID: ownerUser1.ID, Type: 0, Count: 0})
		assert.NoError(err)
		_, err = CreateLog(Log{ItemID: itemID, UserID: user2.ID, OwnerID: ownerUser1.ID, Type: 0, Count: 1})
		assert.NoError(err)
		_, err = CreateLog(Log{ItemID: itemID, UserID: user2.ID, OwnerID: ownerUser2.ID, Type: 0, Count: 0})
		assert.NoError(err)
		_, err = CreateLog(Log{ItemID: itemID, UserID: user1.ID, OwnerID: ownerUser2.ID, Type: 0, Count: 1})
		assert.NoError(err)

		item, err = GetItemByID(item.ID)
		assert.NoError(err)
		log, err := GetLatestLog(item.Logs, ownerUser1.ID)
		assert.NoError(err)
		assert.NotEmpty(log)
		assert.Equal(ownerUser1.Name, log.Owner.Name)
		assert.Equal(user2.ID, log.UserID)
		assert.Equal(user2.Name, log.User.Name)
		assert.Equal(1, log.Count)
	})
}
