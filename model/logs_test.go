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

		log, err = CreateLog(Log{ItemID: 66, OwnerId: 66})
		assert.Error(err)
		assert.Empty(log)
	})

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)

		owner, _ := GetUserByName("traP")
		item, _ := CreateItem(Item{Name: "testItemForCreateLog"})

		log, err := CreateLog(Log{ItemID: item.ID, OwnerId: owner.ID, Type: 0})
		assert.NoError(err)
		assert.NotEmpty(log)
		assert.Equal(owner.ID, log.OwnerId)
		assert.Equal(item.ID, log.ItemID)
	})
}

func TestGetLogsByItemID(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)

		user, err := CreateUser(User{Name: "testGetLogsByItemIDUser"})
		assert.NoError(err)
		owner, err := CreateUser(User{Name: "testGetLogsByItemIDOwner"})
		assert.NoError(err)
		item, err := CreateItem(Item{Name: "testItemForGetLogsByItemID"})
		assert.NoError(err)
		_, err = RegisterOwner(Owner{UserId: owner.ID, Rentalable: true, Count: 1}, item)
		assert.NoError(err)
		log, err := CreateLog(Log{ItemID: item.ID, OwnerId: owner.ID, UserId: user.ID, Type: 0, Count: 0})
		assert.NoError(err)
		logs, err := GetLogsByItemID(item.ID)

		assert.NoError(err)
		assert.NotEmpty(logs)
		assert.Equal(logs[0].ItemID, log.ItemID)
		assert.Equal(logs[0].OwnerId, log.OwnerId)
		assert.Equal(logs[0].Owner.Name, owner.Name)
		assert.Equal(logs[0].User.Name, user.Name)
	})
}

// 今は時間がなくていい感じにテストをかけていませんが未来誰かがテストを書いてくれる日を願って確率でうまくいくテストを残しておきます。テストを書いたらこのコメントは消してください　	ryoha
func TestGetLatestLog(t *testing.T) {

	item, _ := CreateItem(Item{Name: "testGetLatestLogItem"})
	itemID := item.ID
	user, _ := CreateUser(User{Name: "testGetUserLatestLogUser"})
	ownerUser, _ := GetUserByName("traP")
	owner := Owner{
		UserId:     ownerUser.ID,
		Rentalable: true,
		Count:      1,
	}
	_, _ = RegisterOwner(owner, item)
	_, _ = CreateLog(Log{ItemID: itemID, UserId: user.ID, OwnerId: ownerUser.ID, Type: 0, Count: 1})

	t.Run("failures", func(t *testing.T) {
		assert := assert.New(t)

		_, err := RegisterOwner(owner, item)
		assert.NoError(err)
		_, err = CreateLog(Log{ItemID: itemID, UserId: user.ID, OwnerId: ownerUser.ID, Type: 0, Count: 1})
		assert.NoError(err)

		log, err := GetLatestLog(66, 66)
		assert.Error(err)
		assert.Empty(log)

		log, err = GetLatestLog(itemID, 66)
		assert.Error(err)
		assert.Empty(log)
	})

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)

		log, err := GetLatestLog(itemID, ownerUser.ID)
		assert.NoError(err)
		assert.NotEmpty(log)
		t.Log(log)
		assert.Equal(ownerUser.ID, log.OwnerId)
		assert.Equal(ownerUser.Name, log.Owner.Name)
		assert.Equal(user.Name, log.User.Name)
		assert.Equal(itemID, log.ItemID)
		assert.Equal(0, log.Type)
	})
}
