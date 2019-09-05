package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// "fmt"
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

	// failする要素が分らなかったのでfailを書いてません

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		user, _ := CreateUser(User{Name: "registerTestOwner"})
		var owner Owner
		owner.Owner = user
		owner.Rentalable = true
		item, _ := CreateItem(Item{Name: "registerTestItem"})
		item, err := RegisterOwner(owner, item)

		assert.NoError(err)
		assert.NotEmpty(item)
		// fmt.Println(item)
		// assert.Equal(item.Owners[1].Owner.Name, user.Name)
	})
}
