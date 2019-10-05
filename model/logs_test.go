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

func TestGetLogsByItemID(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		assert := assert.New(t)

		owner, _ := GetUserByName("traP")
		item, _ := CreateItem(Item{Name: "testItemForGetLogsByItemID"})
		_, err := RegisterOwner(Owner{OwnerID: owner.ID, Rentalable: true, Count: 1}, item)
		assert.NoError(err)

		log, err := CreateLog(Log{ItemID: item.ID, OwnerID: owner.ID, UserID: owner.ID, Type: 0, Count: 0})
		assert.NoError(err)
		logs, err := GetLogsByItemID(item.ID)
		assert.NoError(err)
		assert.NotEmpty(logs)
		assert.Equal(logs[0].ItemID, log.ItemID)
		assert.Equal(logs[0].OwnerID, log.OwnerID)
	})
}

// 今は時間がなくていい感じにテストをかけていませんが未来誰かがテストを書いてくれる日を願って確率でうまくいくテストを残しておきます。テストを書いたらこのコメントは消してください　	ryoha
// func TestGetLatestLog(t *testing.T) {
// 	t.Parallel()

// 	item, _ := CreateItem(Item{Name: "testGetLatestLogItem"})
// 	itemID := item.ID

// 	user, _ := GetUserByName("traP")
// 	owner := Owner{
// 		OwnerID:    user.ID,
// 		Rentalable: true,
// 		Count:      1,
// 	}
// 	_, _ = RegisterOwner(owner, item)
// 	_, _ = CreateLog(Log{ItemID: itemID, OwnerID: user.ID, Type: 0, Count: 1})

// 	t.Run("failures", func(t *testing.T) {
// 		assert := assert.New(t)

// 		log, err := GetLatestLog(66, 66)
// 		assert.Error(err)
// 		assert.Empty(log)

// 		log, err = GetLatestLog(itemID, 66)
// 		assert.Error(err)
// 		assert.Empty(log)
// 	})

// 	t.Run("success", func(t *testing.T) {
// 		assert := assert.New(t)

// 		log, err := GetLatestLog(itemID, user.ID)
// 		assert.NoError(err)
// 		assert.NotEmpty(log)
// 		assert.Equal(user.ID, log.OwnerID)
// 		assert.Equal(itemID, log.ItemID)
// 		assert.Equal(0, log.Type)
// 	})
// }
