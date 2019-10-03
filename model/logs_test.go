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

		log, err := CreateLog(Log{ItemID: int(item.ID), OwnerID: int(owner.ID)})
		assert.NoError(err)
		assert.NotEmpty(log)
		assert.Equal(int(owner.ID), log.OwnerID)
		assert.Equal(int(item.ID), log.ItemID)
	})
}
