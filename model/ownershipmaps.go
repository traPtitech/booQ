package model

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Ownershipmap ownershipmapsの構造体
type Ownershipmap struct {
	gorm.Model
	ItemID     uint `gorm:"type:int;not null" json:"item_id"`
	Item       Item
	UserID     uint `gorm:"type:int;not null" json:"user_id"`
	User       User
	Rentalable bool `gorm:"default:true" json:"rentalable"`
}

// TableName dbのテーブル名を指定する
func (ownershipmap *Ownershipmap) TableName() string {
	return "ownershipmaps"
}

func CreateOwnershipmap(ownershipmap Ownershipmap) (Ownershipmap, error) {
	if ownershipmap.ItemID == 0 || ownershipmap.UserID == 0 {
		return Ownershipmap{}, errors.New("IDが存在しません")
	}
	db.Create(&ownershipmap)
	return ownershipmap, nil
}
